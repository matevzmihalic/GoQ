package main

import (
	"log"
	"time"
)

func fib(i uint) uint {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	} else {
		return fib(i-1) + fib(i-2)
	}
}

type Fibonacci int

func (w *Fibonacci) Run(number uint, result *uint) error {
	*result = fib(number)
	log.Printf("Input: %d; Output: %d", number, *result)

	if *slow {
		time.Sleep(time.Second)
	}

	return nil
}
