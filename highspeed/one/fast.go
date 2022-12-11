package one

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Server struct{}

type Bread struct{
	name string
	length int32
}

// func init() {
// 	tpl = template.Must(template.ParseFiles("index.html"))
// }

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b := Bread{
		name: "rye",
		length: 12,
	}

	fmt.Fprintf(w, "Welcome to new server on port 8080!")
	parsed, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println("error while rendering: ", err)
	}
	parsed.Execute(w, b.length)
}
