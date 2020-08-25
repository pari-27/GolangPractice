package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() (router *mux.Router) {
	router = mux.NewRouter()

	// No version requirement for /ping
	// router.HandleFunc("/new").Methods(http.MethodGet)
	router.HandleFunc("/", homeHandler).Methods(http.MethodGet)

	return
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/new")
}

var homeTemplate, err = template.ParseFiles("sudoku_interface.html")
