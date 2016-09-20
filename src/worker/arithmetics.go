package main

import (
	"github.com/alfredxing/calc/compute"
)

type Arithmetics int

func (w *Arithmetics) Run(in string, result *float64) error {
	res, err := compute.Evaluate(in)
	if err != nil {
		return err
	}

	*result = res
	return nil
}
