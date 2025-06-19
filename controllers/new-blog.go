package controllers

import (
	"fmt"
	"net/http"
)

func NewBlogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	fmt.Println("data", title, content)
}
