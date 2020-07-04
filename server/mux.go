package main

import (
	"net/http"
)

func newMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Debug(index))
	mux.HandleFunc("/greet", POST(greetHandler))
	// mux.HandleFunc("/greet", POST(greetHandler))

	return mux
}
