package types

import "time"

type Blog struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
