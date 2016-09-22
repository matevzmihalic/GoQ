package main

import (
	"errors"
	"log"
)

type ReverseTextJob struct {
	In     interface{}
	Out    interface{}
	Result chan error
}

var reverseTextQueue (chan ReverseTextJob)
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
	job := ReverseTextJob{in, out, result}
	reverseTextQueue <- job
	return <-result
}

func init() {
	workerReady = map[string](chan bool){
		"ReverseText": make(chan bool),
	}
	reverseTextQueue = make(chan ReverseTextJob)

	go func() {
		for {
			select {
			case job := <-reverseTextQueue:
				go runJob(job, "ReverseText")
			}
		}
	}()
}

func runJob(job ReverseTextJob, workerType string) {
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

func selectWorker(workerType string) *Worker {
	for i, worker := range workers[workerType] {
		if !worker.Busy {
			return &workers[workerType][i]
		}
	}

	<-workerReady[workerType]
	return selectWorker(workerType)
}
