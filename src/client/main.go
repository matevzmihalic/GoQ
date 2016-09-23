package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/rpc"
	"os"
	"time"
)

var serverAddress = flag.String("a", "localhost", "Server address")
var client *rpc.Client

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var aritmeticMethods = []string{"+", "-", "*", "/"}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randomCall() {
	var in interface{}
	var out interface{}
	var method string
	var err error

	switch rand.Intn(3) {
	case 0:
		method = "ReverseText"
		in = randStringRunes(10)
		var result string
		err = client.Call("Q."+method, in, &result)
		out = result
	case 1:
		method = "Arithmetics"
		in = fmt.Sprintf("%d%s%d", rand.Intn(9), aritmeticMethods[rand.Intn(3)], rand.Intn(9))
		var result float64
		err = client.Call("Q."+method, in, &result)
		out = result
	case 2:
		method = "BCrypt"
		in = randStringRunes(10)
		var result string
		err = client.Call("Q."+method, in, &result)
		out = result
	case 3:
		method = "Fibonacci"
		in = uint(rand.Intn(50))
		var result uint
		err = client.Call("Q."+method, in, &result)
		out = result
	}

	if err == nil {
		log.Printf("Called %s(%v) got %v", method, in, out)
	} else {
		log.Printf("Called %s(%v) got ERROR: %v", method, in, err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	flag.Parse()

	var err error
	client, err = rpc.Dial("tcp", *serverAddress+":9001")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		reader.ReadString('\n')
		go randomCall()
	}

}
