package main

import (
	"fmt"
)

type Call struct {
	callerNumber float64
	duration     int
}

type CallCenter struct {
	callsInQueue []Call
	operator     FirstLineSupport
}

func NewCallCenter() *CallCenter {
	return &CallCenter{
		callsInQueue: make([]Call, 0, 10),
		operator:     "lee",
	}
}

func (callCenter CallCenter) addToQueue(call Call) {
	if callCenter.operator.OnLine() {
		callCenter.callsInQueue = append(callCenter.callsInQueue, call)
	}
}

type FirstLineSupport string

func (operator FirstLineSupport) OnLine() bool {
	return true
}

func (operator FirstLineSupport) Away() bool {
	return false
}

func main() {
	fmt.Println("lee is online: ")
}
