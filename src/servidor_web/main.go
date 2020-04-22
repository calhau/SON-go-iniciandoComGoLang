package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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

	//Insert no Banco de dados
	// stmt, err := db.Prepare("Insert into post(title,body) values (?,?)")
	// checkErr(err)

	// _, err = stmt.Exec("Meu segundo posttt", "Meu segundo conteudo")
	// checkErr(err)
	// db.Close()

	rows, err := db.Query("Select * from post")
	checkErr(err)
	// items := []Post{}

	for rows.Next() {
		var id int
		var title string
		var body string

		rows.Scan(&id, &title, &body)
		fmt.Println(id, title, body)
	}

	// Exemplo de rota padrao
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello World")
	})

	// Exemplo de rota com template
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

	//Criando servidor de App na porta 8080
	fmt.Println(http.ListenAndServe(":8080", nil))

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
