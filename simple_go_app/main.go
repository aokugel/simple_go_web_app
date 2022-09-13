package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website of %s", r.URL.Path)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"Birds":[{"Species":"pigeon","Description":"likes to perch on rocks"},{"Species":"eagle","Description":"bird of prey"}],"Number":2}`)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/", apiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
