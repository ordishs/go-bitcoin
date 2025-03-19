package bitcoin

import (
	"context"
	"testing"

	"github.com/libsv/go-bk/bec"
	"github.com/libsv/go-bt/v2"
	"github.com/libsv/go-bt/v2/bscript"
	"github.com/libsv/go-bt/v2/unlocker"
)

func TestSendRawTransactions(t *testing.T) {
	// First we need a private key
	priv, err := bec.NewPrivateKey(bec.S256())
	if err != nil {
		t.Fatal(err)
	}

	address, err := bscript.NewAddressFromPublicKey(priv.PubKey(), false) // false means "not mainnet"
	if err != nil {
		t.Fatal(err)
	}

	lockingScript, err := bscript.NewP2PKHFromAddress(address.AddressString)
	if err != nil {
		t.Fatal(err)
	}

	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	txid, err := b.SendToAddress(address.AddressString, 1.0)
	if err != nil {
		t.Fatal(err)
	}

	fundingTxHex, err := b.GetRawTransactionHex(txid)
	if err != nil {
		t.Fatal(err)
	}

	fundingTx, err := bt.NewTxFromString(*fundingTxHex)
	if err != nil {
		t.Fatal(err)
	}

	var vout int = -1

	// Find the output that we can spend...
	for i, out := range fundingTx.Outputs {
		if out.LockingScript.String() == lockingScript.String() {
			vout = i
			break
		}
	}

	if vout == -1 {
		t.Fatal("Did not find a UTXO")
	}

	tx := bt.NewTx()
	if err := tx.From(fundingTx.TxID(), uint32(vout), fundingTx.Outputs[vout].LockingScriptHexString(), fundingTx.Outputs[vout].Satoshis); err != nil {
		t.Fatal(err)
	}

	amount := fundingTx.Outputs[vout].Satoshis / 2

	if err := tx.PayTo(lockingScript, amount); err != nil {
		t.Fatal(err)
	}

	fq := bt.NewFeeQuote()

	if err := tx.ChangeToAddress(address.AddressString, fq); err != nil {
		t.Fatal(err)
	}

	if err := tx.FillAllInputs(context.Background(), &unlocker.Getter{PrivateKey: priv}); err != nil {
		t.Fatal(err)
	}

	_, err = b.SendRawTransaction(tx.String())
	if err != nil {
		t.Log(err)
	}

	newTxid2, err := b.SendRawTransactionWithoutFeeCheckOrScriptCheck(tx.String())
	if err != nil {
		t.Log(err)
	}

	t.Logf("%+v", newTxid2)
}
