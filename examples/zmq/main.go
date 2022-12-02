package main

import (
	"log"

	"github.com/ordishs/go-bitcoin"
)

func main() {
	zmq := bitcoin.NewZMQ("localhost", 28332)

	ch := make(chan []string)

	go func() {
		for c := range ch {
			log.Printf("%v", c)
		}
	}()

	if err := zmq.Subscribe("hashtx", ch); err != nil {
		log.Fatalf("%v", err)
	}

	if err := zmq.Subscribe("hashblock", ch); err != nil {
		log.Fatalf("%v", err)
	}

	waitCh := make(chan bool)
	<-waitCh
}
