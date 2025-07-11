package main

import (
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
	"personal-blog/controllers"
	"personal-blog/data"
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
	articles, err := controllers.GetArticles()

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	for _, p := range articles {
		logger.Info("Product details",
			slog.Int("id", p.ID),
			slog.String("Title", p.Title),
			slog.String("Content", p.Content),
		)
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/admin.html"))
	tmpl.Execute(w, nil)
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
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/edit.html"))
	tmpl.Execute(w, nil)
}
