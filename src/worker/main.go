package main

import (
	"flag"
	"log"
	"net"
	"net/rpc"
    "os"
    "os/signal"
    "syscall"
)

var workerType = flag.String("t", "Fibonacci", "Worker type")
var workerAddress string

type Worker int

func (w *Worker) SetUp(address string, result *string) error {
	workerAddress = address
	*result = *workerType
	return nil
}

type WorkerAddress struct {
    Type, Address string
}

func cleanup() {
	client, err := rpc.Dial("tcp", ":9001")
    if err != nil {
    	log.Fatal(err)
    }

    var reply int
	client.Call("Control.DisconnectWorker", WorkerAddress{*workerType, workerAddress}, &reply)
}

func main() {
	flag.Parse()

	c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        cleanup()
        os.Exit(1)
    }()

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
