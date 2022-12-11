package payloadserver

import (
	"bufio"
	"fmt"
	"net"
	"reflect"
	"testing"
)

const payload = "The bigger the interface, the weaker the abstraction."

func TestServe(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Error(err)
			return
		}
		defer conn.Close()

		_, err = conn.Write([]byte(payload))
		if err != nil {
			t.Error(err)
		}
		fmt.Println("listening on port 8080, will serve some text ...")
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("connection to port 8080, to get some text ...")
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		t.Error(err)
	}

	expected := []string{"The", "bigger", "the", "interface,", "the", "weaker", "the", "abstraction."}

	if !reflect.DeepEqual(words, expected) {
		t.Fatal("incorrect scanner word list")
	}
	t.Logf("Scanned words: %#v", words)
}
