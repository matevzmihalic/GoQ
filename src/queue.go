package main

import (
	"errors"
	"log"
	"sync"
)

type Job struct {
	In     interface{}
	Out    interface{}
	Result chan error
}

var queue map[string](chan Job)
var workerReady map[string](chan bool)
var busyMutex *sync.Mutex

type WorkerAddress struct {
	Type, Address string
}

type Q int
type Control int

func (c *Control) DisconnectWorker(address WorkerAddress, res *int) error {
	if list, ok := workers[address.Type]; ok {
		busyMutex.Lock()
		for i, worker := range list {
			if worker.Address == address.Address {
				log.Printf("Removing %s worker (%s)\n", address.Type, address.Address)
				workers[address.Type] = append(list[:i], list[i+1:]...)
				busyMutex.Unlock()
				return nil
			}
		}
		busyMutex.Unlock()
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
	busyMutex = &sync.Mutex{}

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

// Runs job on free worker
func runJob(job Job, workerType string) {
	worker := selectWorker(workerType)

	log.Printf("Running %s job (%v) on %s", workerType, job.In, worker.Address)

	job.Result <- worker.Client.Call(workerType+".Run", job.In, job.Out)
	busyMutex.Lock()
	worker.Busy = false
	busyMutex.Unlock()

	select {
	case workerReady[workerType] <- true:
	default:
	}
}

// Selects free worker or waits until one becomes free
func selectWorker(workerType string) *Worker {
	busyMutex.Lock()
	for i, worker := range workers[workerType] {
		if !worker.Busy {
			workers[workerType][i].Busy = true
			busyMutex.Unlock()
			return &workers[workerType][i]
		}
	}
	busyMutex.Unlock()

	<-workerReady[workerType]
	return selectWorker(workerType)
}
