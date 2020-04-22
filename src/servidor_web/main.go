package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

var post = Post{Id: 1, Title: "Meu primeiro Post", Body: "Esse é meu Body"}

// Criando conexao com DB

var db, err = sql.Open("mysql", "root:root@/go_course?charset=utf8")

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/{id}/view", ViewHandler)
	r.HandleFunc("/posts", HomeHandler)
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	//Insert no Banco de dados
	// stmt, err := db.Prepare("Insert into post(title,body) values (?,?)")
	// checkErr(err)

	// _, err = stmt.Exec("Meu segundo posttt", "Meu segundo conteudo")
	// checkErr(err)
	// db.Close()

	// Exemplo de rota padrao
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(rw, "Hello World")
	// })

	// Exemplo de rota com template
	// http.HandleFunc("/template", func(rw http.ResponseWriter, r *http.Request) {
	// 	if title := r.FormValue("title"); title != "" {
	// 		post.Title = title
	// 	}

	// 	t := template.Must(template.ParseFiles("templates/index.html"))
	// 	//t.ExecuteTemplate(rw, "index.html", nil)
	// 	if err := t.ExecuteTemplate(rw, "index.html", post); err != nil {
	// 		http.Error(rw, err.Error(), http.StatusInternalServerError)
	// 	}
	// })

	// Exemplo de rota com template para posts
	// http.HandleFunc("/posts", func(rw http.ResponseWriter, r *http.Request) {
	// 	t := template.Must(template.ParseFiles("templates/posts.html"))
	// 	//t.ExecuteTemplate(rw, "index.html", nil)
	// 	if err := t.ExecuteTemplate(rw, "posts.html", items); err != nil {
	// 		http.Error(rw, err.Error(), http.StatusInternalServerError)
	// 	}
	// })

	//Criando servidor de App na porta 8080
	// fmt.Println(http.ListenAndServe(":8080", nil))
	fmt.Println(http.ListenAndServe(":8080", r))

}

func ListPosts() []Post {
	rows, err := db.Query("Select * from post")
	checkErr(err)
	items := []Post{}

	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Title, &post.Body)
		items = append(items, post)
	}
	return items
}

func GetPostById(id string) Post {
	row := db.QueryRow("select * from post where id=?", id)
	post := Post{}
	row.Scan(&post.Id, &post.Title, &post.Body)
	return post
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/posts.html"))
	//t.ExecuteTemplate(rw, "index.html", nil)
	if err := t.ExecuteTemplate(rw, "posts.html", ListPosts()); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	t := template.Must(template.ParseFiles("templates/view.html"))
	//t.ExecuteTemplate(rw, "index.html", nil)
	if err := t.ExecuteTemplate(rw, "view.html", GetPostById(id)); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
