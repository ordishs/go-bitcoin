package bitcoin

type gbtParams struct {
	Mode         string   `json:"mode,omitempty"`
	Capabilities []string `json:"capabilities,omitempty"`
	Rules        []string `json:"rules,omitempty"`
}

// BlockchainInfo comment
type BlockchainInfo struct {
	Chain                string  `json:"chain"`
	Blocks               int32   `json:"blocks"`
	Headers              int32   `json:"headers"`
	BestBlockHash        string  `json:"bestblockhash"`
	Difficulty           float64 `json:"difficulty"`
	MedianTime           int64   `json:"mediantime"`
	VerificationProgress float64 `json:"verificationprogress,omitempty"`
	Pruned               bool    `json:"pruned"`
	PruneHeight          int32   `json:"pruneheight,omitempty"`
	ChainWork            string  `json:"chainwork,omitempty"`
}

// Transaction comment
type Transaction struct {
	TXID string `json:"txid"`
	Hash string `json:"hash"`
	Data string `json:"data"`
}

// BlockTemplate comment
type BlockTemplate struct {
	Version                  uint32        `json:"version"`
	PreviousBlockHash        string        `json:"previousblockhash"`
	Target                   string        `json:"target"`
	Transactions             []Transaction `json:"transactions"`
	Bits                     string        `json:"bits"`
	CurTime                  uint64        `json:"curtime"`
	CoinbaseValue            uint64        `json:"coinbasevalue"`
	Height                   uint32        `json:"height"`
	MinTime                  uint64        `json:"mintime"`
	NonceRange               string        `json:"noncerange"`
	DefaultWitnessCommitment string        `json:"default_witness_commitment"`
	SizeLimit                uint64        `json:"sizelimit"`
	WeightLimit              uint64        `json:"weightlimit"`
	SigOpLimit               int64         `json:"sigoplimit"`
	VBRequired               int64         `json:"vbrequired"`
	// extra mining candidate fields
	IsMiningCandidate bool     `json:"isminingcandidate"`
	MiningCandidateID string   `json:"miningcandidateid"`
	MerkleBranches    []string `json:"merklebranches"`
}

// MiningCandidate comment
type MiningCandidate struct {
	ID            string   `json:"id"`
	PreviousHash  string   `json:"prevhash"`
	CoinbaseValue uint64   `json:"coinbaseValue"`
	Version       uint32   `json:"version"`
	Bits          string   `json:"nBits"`
	CurTime       uint64   `json:"time"`
	Height        uint32   `json:"height"`
	MerkleProof   []string `json:"merkleProof"`
}

type submitMiningSolutionParams struct {
	ID       string `json:"id"`
	Nonce    uint32 `json:"nonce"`
	Coinbase string `json:"coinbase"`
	Time     uint32 `json:"time"`
	Version  uint32 `json:"version"`
}

// Block comment
type Block struct {
	Hash              string   `json:"hash"`
	Confirmations     int64    `json:"confirmations"`
	Size              uint64   `json:"size"`
	Height            uint32   `json:"height"`
	Version           uint32   `json:"version"`
	Merkleroot        string   `json:"merkleroot"`
	Tx                []string `json:"tx"`
	Time              int64    `json:"time"`
	Nonce             uint64   `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        float64  `json:"difficulty"`
	Chainwork         string   `json:"chainwork,omitempty"`
	Previousblockhash string   `json:"previousblockhash"`
	Nextblockhash     string   `json:"nextblockhash"`
}

// RawTransaction comment
type RawTransaction struct {
	Hex           string `json:"hex"`
	Txid          string `json:"txid"`
	Version       uint32 `json:"version"`
	LockTime      uint32 `json:"locktime"`
	Vin           []Vin  `json:"vin"`
	Vout          []Vout `json:"vout"`
	BlockHash     string `json:"blockhash,omitempty"`
	Confirmations uint64 `json:"confirmations,omitempty"`
	Time          int64  `json:"time,omitempty"`
	Blocktime     int64  `json:"blocktime,omitempty"`
}

// Vout represent an OUT value
type Vout struct {
	Value        float64      `json:"value"`
	N            int          `json:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`
}

// Vin represent an IN value
type Vin struct {
	Coinbase  string    `json:"coinbase"`
	Txid      string    `json:"txid"`
	Vout      int       `json:"vout"`
	ScriptSig ScriptSig `json:"scriptSig"`
	Sequence  uint32    `json:"sequence"`
}

// ScriptPubKey Comment
type ScriptPubKey struct {
	Asm       string   `json:"asm"`
	Hex       string   `json:"hex"`
	ReqSigs   int      `json:"reqSigs,omitempty"`
	Type      string   `json:"type"`
	Addresses []string `json:"addresses,omitempty"`
}

// A ScriptSig represents a scriptsyg
type ScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

// Error comment
type Error struct {
	Code    float64
	Message string
}
