package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type person struct {
	FirstName string
	LastName  string
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Lucy ist da!")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "welcome to the matrix")
}

func me(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "me.gohtml", person{
		FirstName: "Amir",
		LastName:  "Mass",
	})
	if err != nil {
		log.Fatalln(err)
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./me.gohtml"))
}

func main() {

	http.HandleFunc("/", c)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
