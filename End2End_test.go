package bitcoin

import (
	"testing"

	"github.com/bitcoinsv/bsvd/bsvec"
	"github.com/libsv/go-bt"
	"github.com/libsv/go-bt/bscript"
)

func TestSendRawTransactions(t *testing.T) {
	// First we need a private key
	priv, err := bsvec.NewPrivateKey(bsvec.S256())
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
		if out.LockingScript.ToString() == lockingScript.ToString() {
			vout = i
			break
		}
	}

	if vout == -1 {
		t.Fatal("Did not find a UTXO")
	}

	tx := bt.NewTx()
	if err := tx.From(fundingTx.GetTxID(), uint32(vout), fundingTx.Outputs[vout].GetLockingScriptHexString(), fundingTx.Outputs[vout].Satoshis); err != nil {
		t.Fatal(err)
	}

	amount := fundingTx.Outputs[vout].Satoshis / 2

	if err := tx.PayTo(address.AddressString, amount); err != nil {
		t.Fatal(err)
	}

	fees := []*bt.Fee{
		{
			FeeType: "standard",
			MiningFee: bt.FeeUnit{
				Satoshis: 500,
				Bytes:    1000,
			},
			RelayFee: bt.FeeUnit{
				Satoshis: 250,
				Bytes:    1000,
			},
		},
		{
			FeeType: "data",
			MiningFee: bt.FeeUnit{
				Satoshis: 500,
				Bytes:    1000,
			},
			RelayFee: bt.FeeUnit{
				Satoshis: 250,
				Bytes:    1000,
			},
		},
	}

	if err := tx.ChangeToAddress(address.AddressString, fees); err != nil {
		t.Fatal(err)
	}

	if _, err = tx.SignAuto(&bt.InternalSigner{PrivateKey: priv, SigHashFlag: 0}); err != nil {
		t.Fatal(err)
	}

	_, err = b.SendRawTransaction(tx.ToString())
	if err != nil {
		t.Log(err)
	}

	newTxid2, err := b.SendRawTransactionWithoutFeeCheckOrScriptCheck(tx.ToString())
	if err != nil {
		t.Log(err)
	}

	t.Logf("%+v", newTxid2)

}
