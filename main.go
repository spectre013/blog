package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/admin", adminIndex)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("assets/")))
	http.Handle("/", r)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
