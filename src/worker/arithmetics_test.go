package main

import (
	"testing"
)

type arithmeticsTestData struct {
	in  string
	out float64
}

var arithmeticsTests = []arithmeticsTestData{
	{"1+1", 2},
	{"3*3", 9},
	{"3/3", 1},
	{"3+3/3", 4},
}

func TestArithmetics(t *testing.T) {
	reverse := new(Arithmetics)
	for i, test := range arithmeticsTests {
		var out float64
		reverse.Run(test.in, &out)

		if out != test.out {
			t.Errorf("For test case %d (%s, %f)\nexpected: \n\t%f\ngot:\n\t%f", i, test.in, test.out, test.out, out)
		}
	}
}
