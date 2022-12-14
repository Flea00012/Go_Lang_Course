package ticketmachine

import (
	"fmt"
	"math/rand"
	"time"
	"testing"
)

func TestTicketMachine(t *testing.T) {
	var q Queue
	q.New()
	q.StartTicketIssue()
}