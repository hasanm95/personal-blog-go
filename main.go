package main

import (
	"html/template"
	"log"
	"net/http"
	"personal-blog/controllers"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// HTML Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/article/", articleHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/new", newArticleHandler)
	http.HandleFunc("/edit/", editArticleHandler)

	// API Routes
	// http.HandleFunc("/")

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/home.html"))
	tmpl.Execute(w, nil)
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/article.html"))
	tmpl.Execute(w, nil)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/admin.html"))
	tmpl.Execute(w, nil)
}

func newArticleHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/new.html"))
	tmpl.Execute(w, nil)
	controllers.NewBlogHandler(w, r)
}

func editArticleHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/edit.html"))
	tmpl.Execute(w, nil)
}
