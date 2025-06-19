package data

import (
	"encoding/json"
	"fmt"
	"os"
	"personal-blog/types"
)

var blogs []types.Blog

func InitStorage() error {
	file, err := os.Create("posts.json")

	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
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
