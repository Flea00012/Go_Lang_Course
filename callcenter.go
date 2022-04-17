package main

import (
	"fmt"
)

type Call struct {
	callerNumber float64
	duration int
}

type CallCenter struct {
	callsInQueue map[float64] Call
}

type FirstLineSupport struct {
	name string
}

func (operator FirstLineSupport) OnLine() bool {
	return true
}

func (operator FirstLineSupport) away() bool {
	return false
}

func main() {
	var num int
	num = 2
	fmt.Printf("call handled: %d", num)
}
