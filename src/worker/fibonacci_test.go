package main

import (
    "testing"
)

type testData struct {
    in uint
    out uint
}

var tests = []testData {
    {0, 0},
    {1, 1},
    {10, 55},
}

func TestFibonacci(t *testing.T) {
    fibonacci := new(Fibonacci)
    for i, test := range tests {
        var out uint
        fibonacci.Run(test.in, &out)

        if out != test.out {
            t.Errorf("For test case %d (%d, %d)\nexpected: \n\t%d\ngot:\n\t%d", i, test.in, test.out, test.out, out)
        }
    }
}