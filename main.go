package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var atHome bool // false
var num int     // 0
var str string  // ""

func handler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path[1:], "/")

	text := parts[1]
	if text == "" {
		w.WriteHeader(404)
		fmt.Fprint(w, "Failed")
		return
	}

	fmt.Fprintf(w, "Hello, %s", text)
}

func main() {
	http.HandleFunc("/hello/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
