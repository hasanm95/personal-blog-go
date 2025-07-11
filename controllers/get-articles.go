package controllers

import (
	"fmt"
	"personal-blog/data"
	"personal-blog/types"
	"time"
)

func blogViewFromBlog(b types.Blog) types.BlogView {
	return types.BlogView{
		ID:        b.ID,
		Title:     b.Title,
		Content:   b.Content,
		CreatedAt: b.CreatedAt.Format(time.RFC822),
		UpdatedAt: b.UpdatedAt,
	}

}

func GetArticles() ([]types.BlogView, error) {
	blogs, err := data.GetArticles()

	if err != nil {
		return nil, fmt.Errorf("failed to get data %v", err)
	}

	var blogsesView []types.BlogView
	for _, b := range blogs {
		blogsesView = append(blogsesView, blogViewFromBlog(b))
	}

	return blogsesView, nil
}
