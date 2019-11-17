package main

import (
	"fmt"
	"net/http"
)

func homeIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "it work...")
}

func handleRequest() {
	http.HandleFunc("/", homeIndex)
	http.ListenAndServe(":5555", nil)
}

func main() {
	handleRequest()
}
