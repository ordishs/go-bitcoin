package bitcoin

// Force a new commit

import (
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBlockChainInfo(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetBlockchainInfo()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestGetConnectionCount(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	count, err := b.GetConnectionCount()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%d", count)
}

func TestGetNetworkInfo(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetNetworkInfo()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
	t.Logf("Actual IP was %s", b.IPAddress)
}

func TestGetNetTotals(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetNetTotals()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}
func TestMiningInfo(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetMiningInfo()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestUptime(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	count, err := b.Uptime()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%d", count)
}

func TestGetPeerInfo(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetPeerInfo()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestGetRawMempoolWithDetails(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetRawMempool(true)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%s", string(res))
}

func TestGetRawMempoolNoDetails(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetRawMempool(false)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%s", string(res))
}

func TestGetMempoolInfo(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetMempoolInfo()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestGetChainTxStats(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	stats, err := b.GetChainTxStats(0)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", stats)
}

func TestValidateAddress(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	stats, err := b.ValidateAddress("13q47QSaXBHMZVFHFENtTpfd7rtaWTe3v1")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", stats)
}

func TestHelp(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetHelp()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%s", res)
}

func TestGetBestBlockHash(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetBestBlockHash()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%s", res)
}

func TestGetBlockHash(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetBlockHash(1)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%s", res)
}

func TestSendRawTransaction2(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.SendRawTransaction("0100000002b212995825c72d1abf5505a3017694f564dda29908fe5d45ab45b6a3e1748ca5010000006a473044022030d77fbaf3ec092a5038e7fc0235bb2126ca5099dccac5383d58cea01a8d329302207581678aa91743abe31c4cf54e16b0c58e05631d20ccde33be242d96cb57ff1141210373b675c91c95391b6d7977de12bf081bae193470aa0199efc1f847d799497b67ffffffff428947c3e5fb13faf1e71ea86b209b88c3b2e1cd77b326b696fc6e9fb3d0ced6000000006b483045022100801a316894dc40d1e335a14d25b6a7b5b72691b0842878105942d86303a09e1502204d737aecf5d1972e880e329363681bb7c97334e262cf4a20293fee3d8ca87d7041210373b675c91c95391b6d7977de12bf081bae193470aa0199efc1f847d799497b67ffffffff030a1a0000000000001976a914dc9ad4971a54b52308fa3c958df73eac52fb552f88acc6160000000000001976a914a933500e7326f81a974d4212aa16ae29f92e257188ac00000000000000003d6a3b53656e642076696120612057656368617420626f74206d616465206279206161726f6e363720687474703a2f2f6269742e6c792f333331536d754300000000")
	assert.Equal(t, "unexpected response code 500: Missing inputs", err.Error())
	t.Logf("%s", res)
}

func TestSendRawTransaction(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.SendRawTransaction("0100000002b212995825c72d1abf5505a3017694f564dda29908fe5d45ab45b6a3e1748ca5010000006a473044022030d77fbaf3ec092a5038e7fc0235bb2126ca5099dccac5383d58cea01a8d329302207581678aa91743abe31c4cf54e16b0c58e05631d20ccde33be242d96cb57ff1141210373b675c91c95391b6d7977de12bf081bae193470aa0199efc1f847d799497b67ffffffff428947c3e5fb13faf1e71ea86b209b88c3b2e1cd77b326b696fc6e9fb3d0ced6000000006b483045022100801a316894dc40d1e335a14d25b6a7b5b72691b0842878105942d86303a09e1502204d737aecf5d1972e880e329363681bb7c97334e262cf4a20293fee3d8ca87d7041210373b675c91c95391b6d7977de12bf081bae193470aa0199efc1f847d799497b67ffffffff030a1a0000000000001976a914dc9ad4971a54b52308fa3c958df73eac52fb552f88acc6160000000000001976a914a933500e7326f81a974d4212aa16ae29f92e257188ac00000000000000003d6a3b53656e642076696120612057656368617420626f74206d616465206279206161726f6e363720687474703a2f2f6269742e6c792f333331536d754300000000")
	assert.Equal(t, "unexpected response code 500: Missing inputs", err.Error())
	t.Logf("%s", res)
}

func TestSendRawTransactionWithoutFeeCheckOrScriptCheck(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.SendRawTransactionWithoutFeeCheckOrScriptCheck("02000000012cf2bccc35c87b1f6ba2de5e7d175bc0be054230b557efc12636cc092faca6760000000049483045022100f323d358de7e06a43bf1cb272072c283593eaccb90455a6dd5947ae50a0b157b0220212d2208f211ba4be88141cc8c73cf6f9f4116c78e1bb8ffe2d27cb4326ba06f41ffffffff01f0ca052a010000001976a9142a5acfb9a647a03a758afaa5c359284d4b95c0be88ac00000000")
	assert.Equal(t, "Transaction invalid: missing-inputs", err.Error())
	t.Logf("%s", res)
}

func TestGetBlockOverview(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	blk, err := b.GetBlockByHeight(102)
	require.NoError(t, err)

	res, err := b.GetBlockOverview(blk.Hash)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestGetBlock(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	blk, err := b.GetBlockByHeight(102)
	require.NoError(t, err)

	res, err := b.GetBlock(blk.Hash)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestGetBlockByHeight(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetBlockByHeight(102)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestGetBlockStatsByHeight(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetBlockStatsByHeight(102)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestGetBlockStats(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	blk, err := b.GetBlockByHeight(102)
	require.NoError(t, err)

	res, err := b.GetBlockStats(blk.Hash)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestGetGenesisBlock(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	blk, err := b.GetBlockByHeight(0)
	require.NoError(t, err)

	res, err := b.GetBlock(blk.Hash)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestGetBlockHex(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	blk, err := b.GetBlockByHeight(102)
	require.NoError(t, err)

	res, err := b.GetBlockHex(blk.Hash)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%s", *res)
}

func TestGetBlockHeaderHex(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	blk, err := b.GetBlockByHeight(102)
	require.NoError(t, err)

	res, err := b.GetBlockHeaderHex(blk.Hash)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%s", *res)
}

func TestGetBlockHeader(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	blk, err := b.GetBlockByHeight(102)
	require.NoError(t, err)

	res, err := b.GetBlockHeader(blk.Hash)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%+v", res)
}

func TestGetBlockHeaderAndCoinbase(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	blk, err := b.GetBlockByHeight(102)
	require.NoError(t, err)

	res, err := b.GetBlockHeaderAndCoinbase(blk.Hash)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%+v", res)
}

func TestGetRawTransaction(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	blk, err := b.GetBlockByHeight(102)
	require.NoError(t, err)

	res, err := b.GetRawTransaction(blk.Tx[0])
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestGetRawTransactionHex(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	blk, err := b.GetBlockByHeight(102)
	require.NoError(t, err)

	tx, err := b.GetRawTransactionHex(blk.Tx[0])
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", *tx)
}
func TestGetDifficulty(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	diff, err := b.GetDifficulty()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%f", diff)
}

// func TestGetBlockTemplate(t *testing.T) {
// 	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Passing true as an argument will set the segwit rule.  This is necessary for
// 	// BTC and ignored for BCH and BSV.
// 	template, err := b.GetBlockTemplate(true)
// 	if err != nil {
// 		t.Error(err)
// 		t.FailNow()
// 	}
// 	t.Logf("%#v", template)
// }

// func TestGetMiningCandidate(t *testing.T) {
// 	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	template, err := b.GetMiningCandidate()
// 	if err != nil {
// 		t.Error(err)
// 		t.FailNow()
// 	}
// 	t.Logf("%#v", template)
// }

func TestGetSettings(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	template, err := b.GetSettings()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", template)
}

func TestGetTxOut(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	blk, err := b.GetBlockByHeight(102)
	require.NoError(t, err)

	res, err := b.GetTxOut(blk.Tx[0], 0, true)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%#v", res)
}

func TestSubmitBlock(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	template, err := b.SubmitBlock("00000020c803e566bf5af601f19ecb1878f5a6bc2188841c46530c0200000000000000007493d05e5f3ef63975027947bba8c2e08aae8b74e267d9e11e6a4d10e0c0ca3770223e5dc53c081801991f080101000000010000000000000000000000000000000000000000000000000000000000000000ffffffff5f031c0d09046f223e5d2f706f6f6c696e2e636f6d2ffabe6d6ddd5fe0e090d85e6b39d0a6a19bd9b2ff3e3a05113d8ec3bd4a23e39e9c70aec00100000000000000736899c85c811ff6d3eade8b24b8f71213c3bb6d4d003843004000000000ffffffff02807c814a000000001976a91431f2bece272b5d346aa56d09101dd7306d9a307588ac0000000000000000266a24b9e11b6d4ee6b780e5c5ec8a57bc28030aaf0b43632f319648f8bca6abf1ebe3739511f0502c1759")
	assert.Equal(t, "******* BLOCK SUBMIT FAILED with error: <nil> and result: \"inconclusive\"\n", err.Error())
	t.Logf("%#v", template)
}

func TestSubmitMiningSolution(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	var (
		miningCandidateID = "ABC"
		nonce             uint32
		coinbase          string
		time              uint32
		version           uint32
	)

	template, err := b.SubmitMiningSolution(miningCandidateID, nonce, coinbase, time, version)
	assert.Equal(t, "******* BLOCK SUBMIT FAILED with error: unexpected response code 500: bad lexical cast: source type value could not be interpreted as target and result: null\n", err.Error())
	t.Logf("%#v", template)
}

func TestDecodeRawTransactionHex(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := b.DecodeRawTransaction("01000000017235e81ebaccd9fd4d14e3381825c6ac37b56096320e1c54fe00d45e124eca0a0000000000ffffffff01f03a0000000000001976a91424575db8999bc36cd89999de7172b64df2a8893588ac00000000")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", tx)
}

func TestListUnspent(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := b.ListUnspent([]string{"n38vndTAZKFzc3BtPAJ4mecp44UwAZVski"})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for _, utxo := range tx {
		t.Logf("%#v", utxo)
	}
}

func TestSendToAddress(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := b.SendToAddress("n38vndTAZKFzc3BtPAJ4mecp44UwAZVski", 0.01)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(tx)
}

func TestGetRawBlock(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}
	blk, err := b.GetBlockByHeight(102)
	require.NoError(t, err)

	blockBytes, err := b.GetRawBlock(blk.Hash)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(hex.EncodeToString(blockBytes))

	block, err := b.GetBlock(blk.Hash)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%+v", block)
}

func TestGetRawBlockRest(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	blk, err := b.GetBlockByHeight(102)
	require.NoError(t, err)

	r, err := b.GetRawBlockRest(blk.Hash)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(hex.EncodeToString(data))
}

func TestGetRawTransactionRest(t *testing.T) {
	b, err := New("localhost", 18332, "bitcoin", "bitcoin", false)
	if err != nil {
		t.Fatal(err)
	}

	blk, err := b.GetBlockByHeight(102)
	require.NoError(t, err)

	r, err := b.GetRawTransactionRest(blk.Tx[0])
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(hex.EncodeToString(data))
}

func TestSendTime(t *testing.T) {
	keyfunc1 := func(method string, params []interface{}) string {
		var s strings.Builder

		s.WriteString(method)
		s.WriteRune('-')

		s.WriteString(params[0].(string))
		s.WriteRune('|')

		if params[1].(bool) {
			s.WriteRune('T')
		} else {
			s.WriteRune('F')
		}

		if params[2].(bool) {
			s.WriteRune('T')
		} else {
			s.WriteRune('F')
		}

		return s.String()
	}

	keyfunc2 := func(method string, params []interface{}) string {
		return fmt.Sprintf("%s-%v", method, params)
	}

	keyfunc3 := func(method string, params []interface{}) string {
		var s strings.Builder

		fmt.Fprintf(&s, "%s-%s|", method, params[0].(string))

		if params[1].(bool) {
			s.WriteRune('T')
		} else {
			s.WriteRune('F')
		}

		if params[2].(bool) {
			s.WriteRune('T')
		} else {
			s.WriteRune('F')
		}

		return s.String()
	}

	method := "sendrawtransaction"
	hex := "12345677"
	params := []interface{}{hex, false, true}
	times := 1_000_000

	start := time.Now()

	for i := 0; i < times; i++ {
		_ = keyfunc1(method, params)
	}

	t.Logf("1. %s - Took %s\n", keyfunc1(method, params), time.Since(start))

	start = time.Now()

	for i := 0; i < times; i++ {
		_ = keyfunc2(method, params)
	}

	t.Logf("2. %s - Took %s\n", keyfunc2(method, params), time.Since(start))

	start = time.Now()

	for i := 0; i < times; i++ {
		_ = keyfunc3(method, params)
	}

	t.Logf("2. %s - Took %s\n", keyfunc3(method, params), time.Since(start))
}

func TestGetTime(t *testing.T) {
	keyfunc1 := func(method string, params []interface{}) string {
		var b strings.Builder

		b.WriteString(method)
		b.WriteRune('-')
		b.WriteString(params[0].(string))
		b.WriteRune('|')
		b.WriteByte(byte(params[1].(int)))

		return b.String()
	}

	keyfunc2 := func(method string, params []interface{}) string {
		var b strings.Builder

		b.WriteString(method)
		b.WriteRune('-')
		b.WriteString(params[0].(string))
		b.WriteRune('|')
		if params[1].(int) == 0 {
			b.WriteRune('0')
		} else {
			b.WriteRune('1')
		}

		return b.String()
	}

	method := "getrawtransaction"
	hex := "12345677"
	params := []interface{}{hex, 1}
	times := 1_000_000

	start := time.Now()

	for i := 0; i < times; i++ {
		_ = keyfunc1(method, params)
	}

	t.Logf("1. %s - Took %s\n", keyfunc1(method, params), time.Since(start))

	start = time.Now()

	for i := 0; i < times; i++ {
		_ = keyfunc2(method, params)
	}

	t.Logf("2. %s - Took %s\n", keyfunc2(method, params), time.Since(start))

}

func TestDNS(t *testing.T) {
	ips, err := net.LookupIP("google.com")
	if err != nil || len(ips) == 0 {
		t.Error(err)
		t.FailNow()
	}

	for _, i := range ips {
		t.Logf("%t, %#v", i.To4() != nil, i.String())
	}
}

func TestErrTimeout(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "bitcoin", false, WithTimeoutDuration(1*time.Millisecond))
	if err != nil {
		t.Fatal(err)
	}

	_, err = b.ListUnspent([]string{"n38vndTAZKFzc3BtPAJ4mecp44UwAZVski"})
	require.ErrorIs(t, err, ErrTimeout)

}
