package controllers

import (
	"net/http"
	"personal-blog/data"
	"personal-blog/types"
	"strings"
	"time"
)

func UpdateArticle(w http.ResponseWriter, r *http.Request, blog types.Blog) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	title := strings.TrimSpace(r.FormValue("title"))
	content := strings.TrimSpace(r.FormValue("content"))

	if title == "" && content == "" {
		return
	}

	blogs, err := data.GetArticles()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if len(blogs) == 0 {
		blogs = append(blogs, blog)
	}

	for i := range blogs {
		if blogs[i].ID == blog.ID {
			if title != "" {
				blogs[i].Title = title
			}

			if content != "" {
				blogs[i].Content = content
			}

			now := time.Now()
			blogs[i].UpdatedAt = &now

			break
		}
	}
	err = data.RestoreBlog(blogs)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
