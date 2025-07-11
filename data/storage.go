package data

import (
	"encoding/json"
	"fmt"
	"os"
	"personal-blog/types"
)

var blogs []types.Blog

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
