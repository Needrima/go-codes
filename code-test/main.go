package main

import (
	"net/http"
	"html/template"
)

var tpl = template.Must(template.ParseFiles("main.html"))

func foo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl.Execute(w, nil)
	}
}

func Bar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Error occurred", http.StatusMethodNotAllowed)
	}else {
		q := r.FormValue("name")
		tpl.Execute(w, q)
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/search", Bar)

	http.ListenAndServe(":3000", nil)
}