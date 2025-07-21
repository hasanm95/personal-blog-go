package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"personal-blog/controllers"
	"personal-blog/data"
	"personal-blog/types"
)

func main() {
	// Init JSON storage
	err := data.InitStorage()
	if err != nil {
		fmt.Println("%w", err)
	}

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// HTML Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/article/{id}", articleHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/new", newArticleHandler)
	http.HandleFunc("/edit/{id}", editArticleHandler)
	http.HandleFunc("/delete/{id}", deleteArticleHandler)

	// API Routes
	// http.HandleFunc("/")

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	blogs, err := controllers.GetArticles()

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	data := struct {
		Blogs []types.BlogView
	}{
		Blogs: blogs,
	}
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/home.html"))
	tmpl.Execute(w, data)
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	blog := controllers.GetArticleByID(w, r)

	data := struct {
		Blog types.Blog
	}{
		Blog: blog,
	}
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/article.html"))
	tmpl.Execute(w, data)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	blogs, err := controllers.GetArticles()

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	data := struct {
		Blogs []types.BlogView
	}{
		Blogs: blogs,
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/admin.html"))
	tmpl.Execute(w, data)
}

func newArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		controllers.NewBlogHandler(w, r)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/new.html"))
	tmpl.Execute(w, nil)

}

func editArticleHandler(w http.ResponseWriter, r *http.Request) {
	blog := controllers.GetArticleByID(w, r)

	data := struct {
		Blog types.Blog
	}{
		Blog: blog,
	}

	if r.Method == http.MethodPost {
		controllers.UpdateArticle(w, r, blog)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/edit.html"))
	tmpl.Execute(w, data)
}

func deleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	controllers.DeleteArticle(w, r)
}
