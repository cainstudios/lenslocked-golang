package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome to my greatest site!!</h1>\n")
}

func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Println("Starting the server on :3002")
	http.ListenAndServe(":3002", nil)
}
