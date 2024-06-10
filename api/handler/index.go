package handler

import (
	"net/http"
	"text/template"
	
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	template.New("index").ParseFiles()
}
