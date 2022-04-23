package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	bytesReadFrom, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}
	fmt.Print(string(bytesReadFrom))
}
