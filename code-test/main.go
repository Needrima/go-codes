package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type Image struct {
	File multipart.File // file
	Ext  string         // extension
}

func ProcessImage(image Image) {
	bs, err := ioutil.ReadAll(image.File)
	if err != nil {
		log.Println("Byte slice", err)
		return
	}

	tempfile, err := ioutil.TempFile("./images", "*"+image.Ext)
	if err != nil {
		log.Println("tempfile", err)
		return
	}

	tempfile.Write(bs)
}

func Upload(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	if err := os.MkdirAll("./images", 0777); err != nil {
		log.Println("Byte slice", err)
		return
	}

	tpl := template.Must(template.ParseFiles("index.html"))

	var imageNames []string
	if r.Method == http.MethodGet {
		fileinfos, err := ioutil.ReadDir("images")
		if err != nil {
			log.Println("fileinfos", err)
			return
		}

		for _, fileinfo := range fileinfos {
			imageNames = append(imageNames, fileinfo.Name())
		}

		fmt.Println(imageNames)

		tpl.ExecuteTemplate(w, "index.html", imageNames)
	} else if r.Method == http.MethodPost {
		r.ParseMultipartForm(5 << 10) // max of 5MB
		r.ParseForm()

		img1, h1, err := r.FormFile("img1")
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		defer img1.Close()
		imageObject1 := Image{img1, filepath.Ext(h1.Filename)}

		img2, h2, err := r.FormFile("img2")
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		defer img2.Close()
		imageObject2 := Image{img2, filepath.Ext(h2.Filename)}

		img3, h3, err := r.FormFile("img3")
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		defer img3.Close()
		imageObject3 := Image{img3, filepath.Ext(h3.Filename)}

		images := []Image{imageObject1, imageObject2, imageObject3}

		for _, image := range images {
			ProcessImage(image)
		}

		fileinfos, err := ioutil.ReadDir("./images")
		if err != nil {
			log.Println("tempfile", err)
			return
		}

		for _, fileinfo := range fileinfos {
			imageNames = append(imageNames, fileinfo.Name())
		}
		fmt.Println(imageNames)

		tpl.ExecuteTemplate(w, "index.html", imageNames)
	}
}

func main() {
	http.HandleFunc("/", Upload)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Println("Now serving on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
