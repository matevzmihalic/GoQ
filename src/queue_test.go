package main

import (
	"net/rpc"
	"testing"
)

func TestSelectWorker(t *testing.T) {
	go func() {
		workers["ReverseText"] = []Worker{Worker{&rpc.Client{}, "ReverseText", "127.0.0.1:1234", false}}
		workerReady["ReverseText"] <- true
	}()
	selected := selectWorker("ReverseText")

	if selected.Address != "127.0.0.1:1234" {
		t.Errorf("Didn't select worker 127.0.0.1:1235")
	}

	if !selected.Busy {
		t.Errorf("Selected worker not marked as busy")
	}

	workers["ReverseText"] = append(workers["ReverseText"], Worker{&rpc.Client{}, "ReverseText", "127.0.0.1:1235", false})

	selected = selectWorker("ReverseText")
	if selected.Address != "127.0.0.1:1235" {
		t.Errorf("Didn't select worker 127.0.0.1:1235")
	}

	go func() {
		workerReady["ReverseText"] <- true
		workers["ReverseText"][0].Busy = false
		workerReady["ReverseText"] <- true
	}()
	selected = selectWorker("ReverseText")
	if selected.Address != "127.0.0.1:1234" {
		t.Errorf("Didn't select worker 127.0.0.1:1234")
	}
}
