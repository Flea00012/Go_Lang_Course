package main

import (
	"fmt"
	"sync"
)

type Call struct {
	callerNumber float64
	duration     int
}

type CallCenter struct {
	callsInQueue map[float64]Call
	operator     FirstLineSupport
}

func NewCallCenter() *CallCenter {
	return &CallCenter{
		callsInQueue: make(map[string]Call),
		operator: &FirstLineSupport{
			name: "lee",
		}
	}
}

func (callCenter CallCenter) addToQueue(call Call) {
	if callCenter.operator.OnLine() {
		callCenter.callsInQueue = make(call.callerNumber, call)
	}
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
