package main

import (
	"flag"
	"log"
	"net"
	"net/rpc"
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
	case "BCrypt":
		rpc.Register(new(BCrypt))
	case "Arithmetics":
		rpc.Register(new(Arithmetics))
	case "ReverseText":
		rpc.Register(new(ReverseText))
	default:
		*workerType = "Fibonacci"
		rpc.Register(new(Fibonacci))
	}

	conn, err := net.Dial("tcp", ":9002")
	if err != nil {
		log.Fatal(err)
	}

	rpc.ServeConn(conn)
}
