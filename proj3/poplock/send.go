package poplock

import (
	"fmt"
	"net/http"
	"sync"
)

type Messenger struct {
	name string
}

func (m Messenger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("from Messenger: %v ", m.name)
	http.Handle("/", m)
	http.ListenAndServe(":80", m)
}

func (m Messenger) Send(wg *sync.WaitGroup) int {
	start := make(chan struct{})
	n := 1
	foo := func() {
		n++
		defer wg.Done()
		fmt.Println("from foo before go")
		go func() {
			<-start
			fmt.Printf("n is: %d", n)
		}()
		fmt.Println("from foo after go")
	}

	bar := func() {
		defer wg.Done()
		fmt.Println("from bar")
	}

	wg.Add(2)
	go foo()
	go bar()

	close(start)
	wg.Wait()

	return n
}
