package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func dog(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "dog.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func lucyPic(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./lucy.jpg")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./dog.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/lucy.jpg", lucyPic)

	http.ListenAndServe(":8080", nil)
}
