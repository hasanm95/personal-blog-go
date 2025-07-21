package types

import "time"

type Blog struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type BlogView struct {
	ID        int
	Title     string
	Content   string
	CreatedAt string
	UpdatedAt *time.Time
}

type User struct {
	Username string
	Password string
}
