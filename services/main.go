package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":3000", r)
}

// HomeHandler - Home Page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to all about sports\n"))
}
