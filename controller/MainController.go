package controller

import (
	"abel-im/service"
	"html/template"
	"net/http"
)

func Main(writer http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("views/main.html")
	tmp.Execute(writer, nil)
}

func SendMsg(writer http.ResponseWriter, request *http.Request) {
	s := "1"
	service.SendMessage2LocalClient("1", "1", "1", 1, "1", &s)
}
