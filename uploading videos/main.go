package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func uploadVideo(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("file.html"))
	switch r.Method {
	case http.MethodGet:
		tpl.Execute(w, nil)
	case http.MethodPost:
		videoFile, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "No file chosen", http.StatusBadRequest)
			log.Println("Formfile:", err)
			return
		}

		ext := filepath.Ext(header.Filename)
		if ext != ".mp4" && ext != ".mkv" {
			http.Error(w, "File not a video", http.StatusBadRequest)
			return
		}
		fmt.Println(ext)

		bs, err := ioutil.ReadAll(videoFile)
		if err != nil {
			log.Println("ReadAll:", err)
			return
		}

		if err := os.MkdirAll("./videos", 0770); err != nil {
			log.Println("Mkdir:", err)
			return
		}

		tempfile, err := ioutil.TempFile("./videos", "*"+ext)
		if err != nil {
			log.Println("TempFile:", err)
			return
		}
		tempfile.Write(bs)

		var fileNames []string

		fileinfos, err := ioutil.ReadDir("./videos")
		if err != nil {
			log.Println("ReadDir:", err)
			return
		}

		for _, fileinfo := range fileinfos {
			fileNames = append(fileNames, fileinfo.Name())
		}

		tpl.Execute(w, fileNames)

	default:
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", uploadVideo)

	http.Handle("/videos/", http.StripPrefix("/videos/", http.FileServer(http.Dir("videos"))))

	fmt.Println("Serving on port 8080...")
	http.ListenAndServe(":8080", nil)
}
