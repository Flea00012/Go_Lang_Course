package main

import (
	"fmt"
)

type Call struct {
	callerNumber int64
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

func (c CallCenter) HandleCall(operator FirstLineSupport, call Call)  {
	if operator.Away() {
		c.AddToQueue(call)
	} else {
		fmt.Printf("call: %d has been handled by: %s\n", call.callerNumber, c.operator)
	}
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
	call1 := Call{00001, 1}
	call2 := Call{00002, 1}
	call3 := Call{00003, 1}
	callCenter := NewCallCenter()
	fmt.Printf("the calls: %+v to %+v at call-center: %#v\n", call1.callerNumber, lee, callCenter.callsInQueue)

	callCenter.HandleCall(lee, call1)
	callCenter.HandleCall(lee, call2)
	callCenter.HandleCall(lee, call3)
	fmt.Printf("the calls handled: %#v\n", &callCenter.callsInQueue)
}
