package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"web-layout/utils/version"
)

var (
	name   string
	tag    string
	commit string
	branch string
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!\n")
}

func main() {
	var v bool
	flag.BoolVar(&v, "version", false, "show version")
	flag.Parse()

	if v {
		git := version.Info{
			Name:   name,
			Tag:    tag,
			Commit: commit,
			Branch: branch,
		}
		fmt.Println(git.Version())
		return
	}
	http.HandleFunc("/ping", pingHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
