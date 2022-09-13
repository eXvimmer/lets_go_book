package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	// NOTE: you need to Parse flags before using them, otherwise defaults will be used
	flag.Parse()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux := http.NewServeMux()

	// TODO: https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("🚀 Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
