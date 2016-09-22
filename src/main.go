package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	rpc.Register(new(Control))
	rpc.Register(new(Q))

	controlListener, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal(err)
	}
	go rpc.Accept(controlListener)

	listener, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go setUpWorker(conn)
	}
}
