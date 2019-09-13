package bitcoin

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	cache "github.com/patrickmn/go-cache"
	"golang.org/x/sync/singleflight"
)

// A Bitcoind represents a Bitcoind client
type Bitcoind struct {
	client  *rpcClient
	Storage *cache.Cache
	group   singleflight.Group
}

// New return a new bitcoind
func New(host string, port int, user, passwd string, useSSL bool) (*Bitcoind, error) {
	rpcClient, err := newClient(host, port, user, passwd, useSSL)
	if err != nil {
		return nil, err
	}

	defaultExpiration := 5 * time.Second
	cleanupInterval := 10 * time.Second

	return &Bitcoind{
		client:  rpcClient,
		Storage: cache.New(defaultExpiration, cleanupInterval),
		group:   singleflight.Group{},
	}, nil
}

func (b *Bitcoind) call(method string, params []interface{}) (rpcResponse, error) {
	key := fmt.Sprintf("%s|%v", method, params)
	// Check cache
	value, found := b.Storage.Get(key)
	if found {
		// fmt.Printf("CACHED: ")
		return value.(rpcResponse), nil
	}

	// Combine memoized function with a cache store
	value, err, _ := b.group.Do(key, func() (interface{}, error) {
		// fmt.Printf("EXECED: ")
		data, innerErr := b.client.call(method, params)

		if innerErr == nil {
			b.Storage.Set(key, data, cache.DefaultExpiration)
		}

		return data, innerErr
	})
	return value.(rpcResponse), err
}

// GetConnectionCount returns the number of connections to other nodes.
func (b *Bitcoind) GetConnectionCount() (count uint64, err error) {
	r, err := b.call("getconnectioncount", nil)
	if err != nil {
		return 0, err
	}
	count, err = strconv.ParseUint(string(r.Result), 10, 64)
	return
}

// GetBlockchainInfo returns the number of connections to other nodes.
func (b *Bitcoind) GetBlockchainInfo() (info BlockchainInfo, err error) {
	r, err := b.call("getblockchaininfo", nil)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &info)
	return
}

// GetNetworkInfo returns the number of connections to other nodes.
func (b *Bitcoind) GetNetworkInfo() (info NetworkInfo, err error) {
	r, err := b.call("getnetworkinfo", nil)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &info)
	return
}

// GetNetTotals returns the number of connections to other nodes.
func (b *Bitcoind) GetNetTotals() (totals NetTotals, err error) {
	r, err := b.call("getnettotals", nil)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &totals)
	return
}

// GetMiningInfo comment
func (b *Bitcoind) GetMiningInfo() (info MiningInfo, err error) {
	r, err := b.call("getmininginfo", nil)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &info)
	return
}

// Uptime returns the number of connections to other nodes.
func (b *Bitcoind) Uptime() (uptime uint64, err error) {
	r, err := b.call("uptime", nil)
	if err != nil {
		return 0, err
	}
	uptime, err = strconv.ParseUint(string(r.Result), 10, 64)
	return
}

// GetPeerInfo returns the number of connections to other nodes.
func (b *Bitcoind) GetPeerInfo() (info PeerInfo, err error) {
	r, err := b.call("getpeerinfo", nil)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &info)
	return
}

// GetMempoolInfo comment
func (b *Bitcoind) GetMempoolInfo() (info MempoolInfo, err error) {
	r, err := b.call("getmempoolinfo", nil)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &info)
	return
}

// GetRawMempool returns the number of connections to other nodes.
func (b *Bitcoind) GetRawMempool(details bool) (raw []byte, err error) {
	p := []interface{}{details}
	r, err := b.call("getrawmempool", p)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	//err = json.Unmarshal(r.Result, &raw)
	raw, err = json.Marshal(r.Result)
	return
}

// GetChainTxStats returns the number of connections to other nodes.
func (b *Bitcoind) GetChainTxStats(blockcount int) (stats ChainTXStats, err error) {
	p := []interface{}{blockcount}
	r, err := b.call("getchaintxstats", p)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &stats)
	return
}

// ValidateAddress returns the number of connections to other nodes.
func (b *Bitcoind) ValidateAddress(address string) (addr Address, err error) {
	p := []interface{}{address}
	r, err := b.call("validateaddress", p)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &addr)
	return
}

// GetHelp returns the number of connections to other nodes.
func (b *Bitcoind) GetHelp() (j []byte, err error) {
	r, err := b.call("help", nil)
	if err != nil {
		return
	}
	j, err = json.Marshal(r.Result)

	return
}

// GetBestBlockHash comment
func (b *Bitcoind) GetBestBlockHash() (hash string, err error) {
	r, err := b.call("getbestblockhash", nil)
	if err != nil {
		return "", err
	}
	json.Unmarshal(r.Result, &hash)
	return
}

// GetBlockHash comment
func (b *Bitcoind) GetBlockHash(blockHeight int) (blockHash string, err error) {
	p := []interface{}{blockHeight}
	r, err := b.call("getblockhash", p)
	if err != nil {
		return "", err
	}
	json.Unmarshal(r.Result, &blockHash)
	return
}

// SendRawTransaction comment
func (b *Bitcoind) SendRawTransaction(hex string) (txid string, err error) {
	r, err := b.call("sendrawtransaction", []interface{}{hex})
	if err != nil {
		return "", err
	}
	json.Unmarshal(r.Result, &txid)
	return
}

// GetBlock returns information about the block with the given hash.
func (b *Bitcoind) GetBlock(blockHash string) (block *Block, err error) {
	r, err := b.call("getblock", []interface{}{blockHash})

	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &block)
	return
}

// GetBlockOverview returns basic information about the block with the given hash.
func (b *Bitcoind) GetBlockOverview(blockHash string) (block *BlockOverview, err error) {
	r, err := b.call("getblock", []interface{}{blockHash})

	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &block)
	return
}

// GetBlockHex returns information about the block with the given hash.
func (b *Bitcoind) GetBlockHex(blockHash string) (raw *string, err error) {
	r, err := b.call("getblock", []interface{}{blockHash, 0})
	if err != nil {
		return
	}

	err = json.Unmarshal(r.Result, &raw)
	return
}

// GetRawTransaction returns raw transaction representation for given transaction id.
func (b *Bitcoind) GetRawTransaction(txID string) (rawTx *RawTransaction, err error) {
	r, err := b.call("getrawtransaction", []interface{}{txID, 1})
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &rawTx)
	return
}

// GetRawTransactionHex returns raw transaction representation for given transaction id.
func (b *Bitcoind) GetRawTransactionHex(txID string) (rawTx *string, err error) {
	r, err := b.call("getrawtransaction", []interface{}{txID, 0})
	if err != nil {
		return
	}

	err = json.Unmarshal(r.Result, &rawTx)
	return
}

// GetBlockTemplate comment
func (b *Bitcoind) GetBlockTemplate() (template *BlockTemplate, err error) {
	params := gbtParams{
		Mode:         "",
		Capabilities: []string{},
		Rules:        []string{"segwit"},
	}

	r, err := b.call("getblocktemplate", []interface{}{params})
	if err != nil {
		return nil, err
	}

	json.Unmarshal(r.Result, &template)
	return
}

// GetMiningCandidate comment
func (b *Bitcoind) GetMiningCandidate() (template *MiningCandidate, err error) {

	r, err := b.call("getminingcandidate", nil)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(r.Result, &template)
	return
}

// SubmitBlock comment
func (b *Bitcoind) SubmitBlock(hexData string) (result string, err error) {
	r, err := b.client.call("submitblock", []interface{}{hexData})
	if err != nil || r.Err != nil || string(r.Result) != "null" {
		msg := fmt.Sprintf("******* BLOCK SUBMIT FAILED with error: %+v and result: %s\n", err, string(r.Result))
		return "", errors.New(msg)
	}

	return string(r.Result), nil
}

// SubmitMiningSolution comment
func (b *Bitcoind) SubmitMiningSolution(miningCandidateID string, nonce uint32, coinbase string, time uint32, version uint32) (result string, err error) {
	params := submitMiningSolutionParams{
		ID:       miningCandidateID,
		Nonce:    nonce,
		Coinbase: coinbase,
		Time:     time,
		Version:  version,
	}

	r, err := b.client.call("submitminingsolution", []interface{}{params})
	if (err != nil && err.Error() != "") || r.Err != nil || (string(r.Result) != "null" && string(r.Result) != "true") {
		msg := fmt.Sprintf("******* BLOCK SUBMIT FAILED with error: %+v and result: %s\n", err, string(r.Result))
		return "", errors.New(msg)
	}

	return string(r.Result), nil
}

// GetDifficulty comment
func (b *Bitcoind) GetDifficulty() (difficulty float64, err error) {
	r, err := b.call("getdifficulty", nil)
	if err != nil {
		return 0.0, err
	}

	difficulty, err = strconv.ParseFloat(string(r.Result), 64)
	return
}
