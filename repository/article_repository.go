package repository

import (
	"context"
	"database/sql"
	"rest-blog-api/model/domain"
)

type ArticleRepository interface {
	CreateArticle(ctx context.Context, tx *sql.Tx, article domain.Article) domain.Article
	FindAllArticles(ctx context.Context, tx *sql.Tx) []domain.Article
}
