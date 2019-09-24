package main

import (
	"log"

	"github.com/ordishs/go-bitcoin"
)

func main() {
	zmq := bitcoin.NewZMQ("localhost", 28332)

	ch := make(chan string)

	go func() {
		for c := range ch {
			log.Println(c)
		}
	}()

	err := zmq.Subscribe("hashtx", ch)
	if err != nil {
		log.Fatalln(err)
	}

	waitCh := make(chan bool)
	<-waitCh
}
