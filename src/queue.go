package main

import (
	"errors"
	"log"
)

type Job struct {
	In     interface{}
	Out    interface{}
	Result chan error
}

var queue map[string](chan Job)
var workerReady map[string](chan bool)

type WorkerAddress struct {
	Type, Address string
}

type Q int
type Control int

func (c *Control) DisconnectWorker(address WorkerAddress, res *int) error {
	if list, ok := workers[address.Type]; ok {
		for i, worker := range list {
			if worker.Address == address.Address {
				log.Printf("Removing %s worker (%s)\n", address.Type, address.Address)
				workers[address.Type] = append(list[:i], list[i+1:]...)
				return nil
			}
		}
	}
	return errors.New("Worker not found")
}

func (q *Q) ReverseText(in string, out *string) error {
	result := make(chan error)
	queue["ReverseText"] <- Job{in, out, result}
	return <-result
}

func (q *Q) Arithmetics(in string, out *string) error {
	result := make(chan error)
	queue["Arithmetics"] <- Job{in, out, result}
	return <-result
}

func (q *Q) BCrypt(in string, out *string) error {
	result := make(chan error)
	queue["BCrypt"] <- Job{in, out, result}
	return <-result
}

func (q *Q) Fibonacci(in string, out *string) error {
	result := make(chan error)
	queue["Fibonacci"] <- Job{in, out, result}
	return <-result
}

func init() {
	workerReady = map[string](chan bool){
		"ReverseText": make(chan bool),
		"Arithmetics": make(chan bool),
		"BCrypt":      make(chan bool),
		"Fibonacci":   make(chan bool),
	}
	queue = map[string](chan Job){}

	for k := range workerReady {
		queue[k] = make(chan Job)
		go func(k string) {
			for {
				select {
				case job := <-queue[k]:
					go runJob(job, k)
				}
			}
		}(k)
	}
}

func runJob(job Job, workerType string) {
	worker := selectWorker(workerType)

	log.Printf("Running %s job (%v) on %s", workerType, job.In, worker.Address)

	worker.Busy = true
	job.Result <- worker.Client.Call(workerType+".Run", job.In, job.Out)
	worker.Busy = false

	select {
	case workerReady[workerType] <- true:
	default:
	}
}

// Selects free worker or waits until one becomes free
func selectWorker(workerType string) *Worker {
	for i, worker := range workers[workerType] {
		if !worker.Busy {
			return &workers[workerType][i]
		}
	}

	<-workerReady[workerType]
	return selectWorker(workerType)
}
