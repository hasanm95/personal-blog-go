package controllers

import (
	"net/http"
	"personal-blog/data"
	"personal-blog/types"
	"strings"
	"time"

	"github.com/google/uuid"
)

func NewBlogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	title := strings.TrimSpace(r.FormValue("title"))
	content := strings.TrimSpace(r.FormValue("content"))

	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
	}

	if content == "" {
		http.Error(w, "Content is required", http.StatusBadRequest)
	}

	id := int(uuid.New().ID())

	blog := types.Blog{
		ID:        id,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}

	err = data.AddNewBlog(blog)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
