package poplock

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestSender(t *testing.T) {
	fmt.Println("testing sender")
	var wg sync.WaitGroup
	mes := Messenger{name: "sammy"}
	i := mes.Send(&wg)
	if !reflect.DeepEqual(i, 2) {
		fmt.Println("test failed for iterator")
	}
}
