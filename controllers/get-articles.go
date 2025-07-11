package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"personal-blog/types"
)

func GetArticles() ([]types.Blog, error) {
	var articles []types.Blog

	file, err := os.Open("posts.json")

	if err != nil {
		return nil, fmt.Errorf("failed to open file %v", err)
	}

	defer file.Close()

	fileBytes, err := io.ReadAll(file)

	if err != nil {
		return nil, fmt.Errorf("failed to read file %v", err)
	}

	if len(fileBytes) == 0 {
		return articles, nil
	}

	err = json.Unmarshal(fileBytes, &articles)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal file %v", err)
	}

	return articles, nil
}
