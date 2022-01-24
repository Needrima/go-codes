package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

var db = []Person{
	{"Peter", "Parker"},
	{"Barry", "Allen"},
	{"Clark", "Kent"},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		json.NewEncoder(w).Encode(db)
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Accept", "application/json, text/plain, */*")
		w.Header().Set("Content-Type", "x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// var p map[string]interface{}

		// if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		// 	log.Println("DECODER: ", err)
		// 	return
		// }

		// fmt.Println(p)

		// // db = append(db, p)

		// // json.NewEncoder(w).Encode(p)
		
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
