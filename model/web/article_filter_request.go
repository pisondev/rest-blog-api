package web

import "time"

type ArticleFilterRequest struct {
	Title     string
	StartDate time.Time
	EndDate   time.Time
}
