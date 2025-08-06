package domain

import "time"

type ArticleFilter struct {
	Title     string
	StartDate time.Time
	EndDate   time.Time
}
