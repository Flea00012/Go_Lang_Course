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

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Panic(err)
		return
	}
	defer file.Close()

	var (
		b = make([]byte, 16)
	)
	for n := 0; err == nil; {
		n, err = file.Read(b)
		if err == nil {
			fmt.Print(string(b[:n]))
		}
	}
	if err != nil && err != io.EOF {
		fmt.Println("\n\nError:", err)
	}
}
