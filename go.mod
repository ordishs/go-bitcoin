module github.com/ordishs/go-bitcoin

go 1.13

replace (
	bitbucket.org/nchteamnch/common v1.0.3 => ../common
	bitbucket.org/simon_ordish/cryptolib v1.0.33 => ../cryptolib
	github.com/ordishs/go-bitcoin v1.0.16 => ../go-bitcoin
	github.com/ordishs/gocore v1.0.7 => ../gocore
)

require (
	bitbucket.org/simon_ordish/cryptolib v1.0.33
	github.com/go-zeromq/zmq4 v0.5.0
	github.com/ordishs/gocore v1.0.7
	github.com/patrickmn/go-cache v2.1.0+incompatible
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
)
