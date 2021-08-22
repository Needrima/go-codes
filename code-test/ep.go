package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type datum struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Age   int    `json:"age"`
}

var data = map[string]datum{
	"spiderman": {"Peter", "Parker", 18},
	"superman":  {"Clark", "Kent", 30},
	"batman":    {"Bruce", "Wayne", 30},
}

func foo(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	key = strings.ToLower(key)
	fmt.Println(key)
	
	w.Header().Set("Content-type", "application/json")

	datum, found := data[key]
	if !found {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`No data for key`))
		return
	}

	err := json.NewEncoder(w).Encode(datum)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}
