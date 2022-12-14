package ticketmachine

import (
	"fmt"
)

const(
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

type Queue struct {
	waitPass int
	waitTicket int
	playPass bool
	playTicket bool
	queuePass chan int
	queueTicket chan int
	message chan int
}

func (q *Queue) New()  {
	q.queuePass = make(chan int)
	q.queueTicket = make(chan int)
	q.message = make(chan int)

	go func ()  {
		var message int
		for {
			select {
			case message = <- q.message:
				switch message {
				case messagePassStart:
					q.waitPass++
				case messagePassEnd:
					q.playPass = false
				case messageTicketStart:
					q.waitTicket++
				case messageTicketEnd:
					q.playTicket = false	
				}
				if q.waitPass > 0 && q.waitTicket > 0 && !q.playPass &&!q.playTicket {
					q.playPass = true
					q.playTicket = true
					q.waitTicket--
					q.waitPass--
					q.queuePass <- 1
					q.queueTicket <- 1
				}
			}
		}
	}()
}

func (q *Queue) StartTicketIssue()  {
	q.message <- messageTicketStart
	<- q.queueTicket
}

func main() {
	
}