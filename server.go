package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const port = ":8080"

type News struct {
	Headline string
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	tpl := template.Must(template.ParseFiles("./static/hangman.gohtml"))
	tpl.Execute(w, nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	event := News{Headline: " "}
	tpl := template.Must(template.ParseFiles("./static/index.gohtml"))
	tpl.Execute(w, event)
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("Extra/"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hangman", formHandler)
	http.HandleFunc("/index", helloHandler)
	http.Handle("/Extra/", http.StripPrefix("/Extra/", http.FileServer(http.Dir("Extra"))))

	//http.Handle("/", http.FileServer(http.Dir("./hangman")))
	//http.ListenAndServe(":8080", nil)
	//fmt.Println("(http://localhost:8080/hangman) - Server started on port", port)
	//http.ListenAndServe(port, nil)

	fmt.Printf("http://localhost:8080/hangman - Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
