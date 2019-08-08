package bitcoin

import "testing"

func TestGetBlockChainInfo(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
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
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
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

func TestGetBlock(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetBlock("000000000000000005e827eecfc1b8cbb990f4ae458e748480d10b80458faf25")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestGetRawTransaction(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := b.GetRawTransaction("4dd023de1efdd78129780da35ba35411868480fc2240ddd7ba042823ae02cdc9")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", tx)
}

func TestGetRawTransactionHex(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := b.GetRawTransactionHex("4dd023de1efdd78129780da35ba35411868480fc2240ddd7ba042823ae02cdc9")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", tx)
}

func TestGetBlockTemplate(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	template, err := b.GetBlockTemplate()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", template)
}

func TestGetMiningCandidate(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	template, err := b.GetMiningCandidate()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", template)
}

func TestSubmitBlock(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	template, err := b.SubmitBlock("00000020c803e566bf5af601f19ecb1878f5a6bc2188841c46530c0200000000000000007493d05e5f3ef63975027947bba8c2e08aae8b74e267d9e11e6a4d10e0c0ca3770223e5dc53c081801991f080101000000010000000000000000000000000000000000000000000000000000000000000000ffffffff5f031c0d09046f223e5d2f706f6f6c696e2e636f6d2ffabe6d6ddd5fe0e090d85e6b39d0a6a19bd9b2ff3e3a05113d8ec3bd4a23e39e9c70aec00100000000000000736899c85c811ff6d3eade8b24b8f71213c3bb6d4d003843004000000000ffffffff02807c814a000000001976a91431f2bece272b5d346aa56d09101dd7306d9a307588ac0000000000000000266a24b9e11b6d4ee6b780e5c5ec8a57bc28030aaf0b43632f319648f8bca6abf1ebe3739511f0502c1759")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", template)
}

func TestSubmitMiningSolution(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
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
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", template)
}
