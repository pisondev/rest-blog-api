package domain

import "time"

type Article struct {
	Id        int
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
