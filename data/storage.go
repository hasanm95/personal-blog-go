package data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"personal-blog/types"
)

func InitStorage() error {
	fileName := "posts.json"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {

		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	return nil
}

func AddNewBlog(newBlog types.Blog) error {
	blogs, err := GetArticles()

	if err != nil {
		return fmt.Errorf("failed to get data: %w", err)
	}

	blogs = append(blogs, newBlog)
	blogBytes, err := json.MarshalIndent(blogs, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	err = os.WriteFile("posts.json", blogBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write data to json: %w", err)
	}

	return nil
}

func GetArticles() ([]types.Blog, error) {
	var blogs []types.Blog
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
		return blogs, nil
	}

	err = json.Unmarshal(fileBytes, &blogs)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal file %v", err)
	}

	return blogs, nil
}
