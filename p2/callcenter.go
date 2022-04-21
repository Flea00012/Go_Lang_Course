package main

import (
	"fmt"
)

type Call struct {
	callerNumber float64
	duration     int64
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

func (c CallCenter) HandleCall(operator FirstLineSupport, call Call) (Call, []Call) {
	if operator.OnLine() {
		c.AddToQueue(call)
	}
	return call, c.callsInQueue
}

func (callCenter CallCenter) AddToQueue(call Call) {
	callCenter.callsInQueue = append(callCenter.callsInQueue, call)
}

type FirstLineSupport string

func (operator FirstLineSupport) OnLine() bool {
	return true
}

func (operator FirstLineSupport) Away() bool {
	return false
}

func main() {
	var lee FirstLineSupport
	lee = "lee"
	call := Call{00001, 1}
	callCenter := NewCallCenter()
	fmt.Printf("the calls: %T to %T at call-center: %T\n", call.callerNumber, lee, callCenter.callsInQueue)

	_, calls := callCenter.HandleCall(lee, call)
	fmt.Printf("the calls handled: %T\n", calls)
}
