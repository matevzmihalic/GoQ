package main

type ReverseText int

func (w *ReverseText) Run(in string, out *string) error {
    runes := []rune(in)
    length := len(runes)
    for i := 0; i < length/2; i++ { 
        runes[i], runes[length-1-i] = runes[length-1-i], runes[i] 
    }
    *out = string(runes)
    return nil
}