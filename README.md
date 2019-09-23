# go-bitcoin
Go wrapper for bitcoin RPC

Start by creating a connection to a bitcoin node
```
  b, err := New("rcp host", rpc port, "rpc username", "rpc password", false)
  if err != nil {
    log.Fatal(err)
  }
```

Then make a call to bitcoin
```
  res, err := b.GetBlockchainInfo()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%#v\n", res)
```

Available calls are:
```
GetConnectionCount()
GetBlockchainInfo()
GetNetworkInfo()
GetNetTotals()
GetMiningInfo()
Uptime()
GetMempoolInfo()
GetRawMempool(details bool)
GetChainTxStats(blockcount int)
ValidateAddress(address string)
GetHelp()
GetBestBlockHash()
GetBlockHash(blockHeight int)
SendRawTransaction(hex string)
GetBlock(blockHash string)
GetBlockOverview(blockHash string)
GetBlockHex(blockHash string)
GetRawTransaction(txID string)
GetRawTransactionHex(txID string)
GetBlockTemplate(includeSegwit bool)
GetMiningCandidate()
SubmitBlock(hexData string)
SubmitMiningSolution(candidateID string, nonce uint32,
                     coinbase string, time uint32, version uint32)
GetDifficulty()


```