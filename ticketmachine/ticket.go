package ticketmachine

import (
	"fmt"
)

const(
	messageStartPass = iota
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
				case messageStartPass:
					
				}
			}

		}
	}()

}

func main() {
	
}