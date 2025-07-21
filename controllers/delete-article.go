package controllers

import (
	"net/http"
	"personal-blog/data"
	"strconv"
)

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	blogId := r.PathValue("id")

	if blogId == "" {
		http.Error(w, "provide a blog id", http.StatusBadRequest)
	}

	intId, err := strconv.Atoi(blogId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	blogs, err := data.GetArticles()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var blogToDelete int

	for i := range blogs {
		if blogs[i].ID == intId {
			blogToDelete = i
			break
		}
	}

	blogs = append(blogs[:blogToDelete], blogs[blogToDelete+1:]...)
	err = data.RestoreBlog(blogs)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
