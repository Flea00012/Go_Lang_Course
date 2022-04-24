package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify a path.")
		return
	}

	// b, err := ioutil.ReadFile(os.Args[1])

	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(string(b))

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	var (
		b = make([]byte, 16)
	)
	for n := 0; err == nil; {
		n, err := file.Read(b)
		if err == nil {
			fmt.Println(string(b[:n]))
		}
	}
	if err != nil && err != io.EOF {
		fmt.Println("\n\nError:", err)
	}
}
