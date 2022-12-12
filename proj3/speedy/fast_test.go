package speedy

import (
	"fmt"
	// "reflect"
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

func TestChannel(t *testing.T) {
	fmt.Println("testing of the producer routines")
	ch := make(chan int, 10)
	defer close(ch)
	wg.Add(2)
	go producerOne(ch)
	go producerTwo(ch)

	fmt.Printf("chan values: %#v\n", ch)

	wg.Wait()
	// expected := [10]int {0,1,2,3,4,10,9,8,7,6}
	// if !reflect.DeepEqual(ch, expected) {
	// 	panic("Test for channel failed")
	// }
	
}
