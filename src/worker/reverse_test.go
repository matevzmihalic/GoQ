package main

import (
    "testing"
)

type reverseTestData struct {
    in string
    out string
}

var reverseTests = []reverseTestData {
    {"abc", "cba"},
    {"čćžšđ", "đšžćč"},
    {"世界", "界世"},
}

func TestReverseText(t *testing.T) {
    reverse := new(ReverseText)
    for i, test := range reverseTests {
        var out string
        reverse.Run(test.in, &out)

        if out != test.out {
            t.Errorf("For test case %d (%s, %s)\nexpected: \n\t%s\ngot:\n\t%s", i, test.in, test.out, test.out, out)
        }
    }
}