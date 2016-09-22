package main

import (
	"log"
	"time"
)

type ReverseText int

func (w *ReverseText) Run(in string, out *string) error {
	runes := []rune(in)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	*out = string(runes)
	log.Printf("Input: %s; Output: %s", in, *out)

	if *slow {
		time.Sleep(time.Second)
	}

	return nil
}
