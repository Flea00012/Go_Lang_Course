package main

import (
	"bufio"
	"fmt"
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
		go handleConnection(acc)
	}

}

func handleConnection(acc net.Conn) {
	err := acc.SetDeadline(time.Now().Add(10*time.Second))
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
