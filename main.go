package main

import (
	"fmt"
	"golang-web/lib"
	"html/template"
	"log"
	"net/http"
	"os"
)

var Hostname string

type Page struct {
	Title string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	p := &Page{Title: lib.IndexTitle(Hostname)}

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, p)
}

func main() {
	// http
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	Hostname, _ = os.Hostname()

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
