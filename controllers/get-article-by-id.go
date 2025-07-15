package controllers

import (
	"net/http"
	"personal-blog/data"
	"personal-blog/types"
	"strconv"
)

func GetArticleByID(w http.ResponseWriter, r *http.Request) types.Blog {
	blogId := r.PathValue("id")

	if blogId == "" {
		http.Error(w, "provide a blog id", http.StatusBadRequest)
		return types.Blog{}
	}

	intId, err := strconv.Atoi(blogId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	blogs, err := data.GetArticles()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return types.Blog{}
	}

	if len(blogs) == 0 {
		http.Error(w, "no blog found", http.StatusBadRequest)
		return types.Blog{}
	}

	var blog types.Blog

	for i := range blogs {
		if blogs[i].ID == intId {
			blog = blogs[i]
			break
		}
	}

	if blog == (types.Blog{}) {
		http.Error(w, "no blog found with id", http.StatusBadRequest)
		return types.Blog{}
	}

	return blog
}
