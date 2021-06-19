//sorting strings
package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"os"
	"fmt"
)

var gridFs *mgo.GridFS
var filename string

func StoreImage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		template.Must(template.ParseFiles("index.html")).Execute(w, nil)
	} else if r.Method == http.MethodPost {
		file, header, err := r.FormFile("img")
		if err != nil {
			log.Fatalln("file", err)
			return
		}
		defer file.Close()
		filename = header.Filename
		log.Println("filename", filename)		

		bs, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatalln("Readall", err)
			return
		}

		filedb, err := gridFs.Create(filename)
		defer filedb.Close()
		if err != nil {
			log.Fatalln("Gridfs file", err)
			return
		}

		if _, err = filedb.Write(bs); err != nil {
			log.Println("write", err)
		}

		template.Must(template.ParseFiles("index.html")).Execute(w, "Done")
	}

}

func RetrieveImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Unaccepted method", http.StatusMethodNotAllowed)
		return
	}

	storage := make([]byte, 8192)

	if err := gridFs.Find(bson.M{"filename": filename}).One(&storage); err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if _, err := os.Create(fmt.Sprintf("%s%s", storage,filepath.Ext(filename))); err != nil {
		log.Println(err)
	}


	template.Must(template.ParseFiles("gallery.html")).Execute(w, nil)
}

func main() {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatalln("session:", err)
		return
	}
	defer session.Close()

	gridFs = session.DB("golang").GridFS("go_practice")

	http.HandleFunc("/", StoreImage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
