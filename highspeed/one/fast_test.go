package one

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

// type Server struct {}

// func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to new server on port 8080!")
// }

func TestServers(t *testing.T)  {
	http.Handle("/", Server{})
	s := http.TimeoutHandler(Server{},time.Second*5, "Handler  was alive for more than 5 seconds")
	http.ListenAndServe(":8080", s)	
}

func TestGoFast(t *testing.T) {
	fmt.Println("lets go faster")
	s := []byte("This is the string to write to file. Its so much from.")

	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	ioutil.WriteFile("file.txt", s, 0644)

}


