package speedy

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup
var sharedLock sync.Mutex
var num int32

func producerOne(ch chan int) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		num--
		fmt.Println("producerOne run: ", i)
		sharedLock.Unlock()
		ch <- i
	}

}

func producerTwo(ch chan int) {
	defer wg.Done()

	for i := 10; i > 5; i-- {
		sharedLock.Lock()
		time.Sleep(2 * time.Nanosecond)
		num++
		fmt.Println("producerTwo run: ", num)
		sharedLock.Unlock()
		ch <- i
	}

}

func TestXxx(t *testing.T) {
	fmt.Println("testing of the producer routines")
	ch := make(chan int, 10)
	wg.Add(2)
	go producerOne(ch)
	go producerTwo(ch)

	fmt.Printf("chan values: %#v\n", ch)

	wg.Wait()
	close(ch)
}
