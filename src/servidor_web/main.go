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

var post = Post{Id: 1, Title: "Meu primeiro Post", Body: "Esse é meu Body"}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello World")
	})

	http.HandleFunc("/template", func(rw http.ResponseWriter, r *http.Request) {
		if title := r.FormValue("title"); title != "" {
			post.Title = title
		}

		t := template.Must(template.ParseFiles("templates/index.html"))
		//t.ExecuteTemplate(rw, "index.html", nil)
		if err := t.ExecuteTemplate(rw, "index.html", post); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))

}
