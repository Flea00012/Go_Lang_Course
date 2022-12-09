package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

type serv string

func (s serv) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("youve got mail")
}

// func main() {
// 	err := tpl.ExecuteTemplate(os.Stdout, "index.html", "ons is lekker")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }
