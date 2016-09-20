package main

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
	return nil
}
