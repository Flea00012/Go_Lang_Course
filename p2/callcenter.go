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
	if operator.Away() {
		callPickedUp := c.callsInQueue[len(c.callsInQueue)-1]
		newQueue := make([]Call, len(c.callsInQueue)-1)
		newQueue2 := append(newQueue, c.callsInQueue[len(c.callsInQueue)-1])
		// fmt.Println("%s handled the call from %s", operator, call.callerNumber)
		return callPickedUp, newQueue2
	}
	fmt.Println("call from s% not handled and added to Queue", call.callerNumber)
	c.AddToQueue(call)
	return call, c.callsInQueue
}

func (callCenter CallCenter) AddToQueue(call Call) {
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
	// fmt.Println("lee is online: ")
	var lee FirstLineSupport
	lee = "lee"
	call := Call{00001, 1}
	callCenter := NewCallCenter()
	fmt.Printf("the calls: %T to %T at center: %T", call, lee, callCenter)
	
	call, calls := callCenter.HandleCall(lee, call)
	fmt.Printf("the calls handled: %T", calls)
}
