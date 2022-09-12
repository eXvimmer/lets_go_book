package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("View a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// NOTE: servemux in Go terminology = router
	mux := http.NewServeMux() // a new ServeMux: an HTTP request multiplexer

	// NOTE: subtree path (ends with a trailing / e.g. /** or /static/** ->
	// NOTE: acts like a catch-all)
	mux.HandleFunc("/", home)

	// NOTE: fixed paths (not ending with /): request URL path should exactly
	// NOTE: match the fixed path
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// TODO: make a variable for PORT
	log.Println("🚀 Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
