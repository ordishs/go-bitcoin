package bitcoin

import (
	"encoding/hex"
	"testing"
)

func TestCoinbase(t *testing.T) {
	h, _ := hex.DecodeString("01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff580320a107152f5669614254432f48656c6c6f20576f726c64212f2cfabe6d6dbcbb1b0222e1aeebaca2a9c905bb23a3ad0302898ec600a9033a87ec1645a446010000000000000010f829ba0b13a84def80c389cde9840000ffffffff0174fdaf4a000000001976a914f1c075a01882ae0972f95d3a4177c86c852b7d9188ac00000000")

	tx, _ := TransactionFromBytes(h)

	expected := "5ebaa53d24c8246c439ccd9f142cbe93fc59582e7013733954120e9baab201df"

	if tx.Hash != expected {
		t.Errorf("Expected %q, got %q", expected, tx.Hash)
	}
}

func TestToHex(t *testing.T) {
	hexStr := "01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff580320a107152f5669614254432f48656c6c6f20576f726c64212f2cfabe6d6dbcbb1b0222e1aeebaca2a9c905bb23a3ad0302898ec600a9033a87ec1645a446010000000000000010f829ba0b13a84def80c389cde9840000ffffffff0174fdaf4a000000001976a914f1c075a01882ae0972f95d3a4177c86c852b7d9188ac00000000"
	h, _ := hex.DecodeString(hexStr)

	tx, _ := TransactionFromBytes(h)

	res := hex.EncodeToString(tx.ToHex())

	if res != hexStr {
		t.Errorf("Expected %q, got %q", hexStr, res)
	}
}
