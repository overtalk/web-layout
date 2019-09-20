package main

import (
	"fmt"
	"log"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!\n")
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
