package main

import (
	"github.com/alfredxing/calc/compute"
	"log"
	"time"
)

type Arithmetics int

func (w *Arithmetics) Run(in string, result *float64) error {
	res, err := compute.Evaluate(in)
	if err != nil {
		return err
	}

	*result = res
	log.Printf("Input: %s; Output: %f", in, *result)

	if *slow {
		time.Sleep(time.Second)
	}

	return nil
}
