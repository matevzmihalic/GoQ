package main

import (
	"log"
	"net"
	"net/rpc"
)

type Worker struct {
	Client  *rpc.Client
	Type    string
	Address string
	Busy    bool
}

var workers map[string][]Worker

func setUpWorker(conn net.Conn) {
	client := rpc.NewClient(conn)
	var workerType string

	if err := client.Call("Worker.SetUp", conn.RemoteAddr().String(), &workerType); err != nil {
		log.Printf("Couldn't set up worker: %v\n", err)
		return
	}

	log.Printf("Connected %s worker (%v)", workerType, conn.RemoteAddr())

	worker := Worker{client, workerType, conn.RemoteAddr().String(), false}
	if list, ok := workers[workerType]; ok {
		workers[workerType] = append(list, worker)
	} else {
		workers[workerType] = []Worker{worker}
	}

	select {
	case workerReady[workerType] <- true:
	default:
	}

}

func init() {
	workers = make(map[string][]Worker)
}
