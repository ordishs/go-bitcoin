package main

import (
	"github.com/ordishs/go-bitcoin"
	"github.com/ordishs/gocore"
)

func main() {
	zmq := bitcoin.NewZMQ("localhost", 28332)

	ch := make(chan []string)

	go func() {
		for c := range ch {
			logger.Infof("%v", c)
		}
	}()

	if err := zmq.Subscribe("hashtx", ch); err != nil {
		logger.Fatalf("%v", err)
	}

	if err := zmq.Subscribe("hashblock", ch); err != nil {
		logger.Fatalf("%v", err)
	}

	waitCh := make(chan bool)
	<-waitCh
}
