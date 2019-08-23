package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/upload", handleUpload)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("sertver started at locahost:9900")
	http.ListenAndServe(":9000", nil)

}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "only accept post request", http.StatusBadRequest)
		return
	}
	basePath, _ := os.Getwd()
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		fileLocation := filepath.Join(basePath, "files", part.FileName())
		dst, err := os.Create(fileLocation)
		if dst != nil {
			defer dst.Close()
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if _, err := io.Copy(dst, part); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
	w.Write([]byte("all file uploaded"))

}
