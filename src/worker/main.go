package main

import (
	"net"
	"net/rpc"
	"log"
	"flag"
)

var workerType = flag.String("t", "Fibonacci", "Worker type")

type Worker int

func (w *Worker) GetType(empty int, result *string) error {
	*result = *workerType
	return nil
}

func main() {
	flag.Parse()

	baseWorker := new(Worker)
	rpc.Register(baseWorker)

	switch *workerType {
	case "ReverseText":
		rpc.Register(new(ReverseText))
	default:
		*workerType = "Fibonacci"
		rpc.Register(new(Fibonacci))
	}

	conn,err := net.Dial("tcp",":9002")
	if err != nil {
		log.Fatal(err)
	}

	rpc.ServeConn(conn)
}