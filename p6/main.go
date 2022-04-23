package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		acc, err := li.Accept()
		if err != nil {
			log.Printf("error: %e", err)
		}
	
		io.WriteString(acc, "hello from outside lambda\n")
		go handleConnection(acc)
		go func() {
			io.WriteString(acc, "hello from inside lambda\n")
		}()
	}

}

func handleConnection(acc net.Conn) {
	err := acc.SetDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		log.Println("Connection timeout")
	}
	scanner := bufio.NewScanner(acc)
	for scanner.Scan() {
		readLine := scanner.Text()
		fmt.Println(readLine)
		fmt.Fprintln(acc, "lets start chatting", readLine)
	}
	defer acc.Close()
}
