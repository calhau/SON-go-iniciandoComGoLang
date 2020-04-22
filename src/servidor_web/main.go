package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello World")
	})

	http.HandleFunc("/template", func(rw http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("templates/index.html"))
		//t.ExecuteTemplate(rw, "index.html", nil)
		if err := t.ExecuteTemplate(rw, "index.html", nil); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))

}
