package controller

import (
	"html/template"
	"net/http"
)

func Main(writer http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("views/main.html")
	tmp.Execute(writer, nil)
}

