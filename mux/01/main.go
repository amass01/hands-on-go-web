package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Lucy ist da!")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "welcome to the matrix")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "amass is here")
}

func main() {

	http.HandleFunc("/", c)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
