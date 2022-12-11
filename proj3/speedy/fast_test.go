package speedy

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup
var sharedLock sync.Mutex
var ch chan int

func producerOne(ch chan int) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		fmt.Println("producerOne run: ", i)
		sharedLock.Unlock()
	}

}

func producerTwo(ch chan int) {
	defer wg.Done()

	for i := 10; i < 0; i-- {
		sharedLock.Lock()
		time.Sleep(2 * time.Nanosecond)
		fmt.Println("producerOne run: ", i)
		
		sharedLock.Unlock()
	}

}

func TestXxx(t *testing.T) {
	fmt.Println("testing of the producer routines")

	wg.Add(2)
	go producerOne(ch)
	go producerTwo(ch)

	wg.Wait()

}
