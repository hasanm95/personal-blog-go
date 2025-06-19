package data

import (
	"fmt"
	"os"
)

func InitStorage() error {
	file, err := os.Create("posts.json")

	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer file.Close()

	return nil
}
