package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplPath := filepath.Join("templates", "index.html")
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			log.Println("Template parse error: ", err)
			return
		}
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/cv", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/cv.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			log.Println("Template error:", err)
			return
		}
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/contact.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			log.Println("Template error:", err)
			return
		}
		tmpl.Execute(w, nil)
	})

	log.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
