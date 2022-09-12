package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox")) // typecast string to []byte
}

func main() {
	// NOTE: servemux in Go terminology = router
	mux := http.NewServeMux() // a new ServeMux: an HTTP request multiplexer
	mux.HandleFunc("/", home)

	// TODO: make a variable for PORT
	log.Println("🚀 Starting server on :4000")

	// NOTE: the TCP network address format: host:port
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
