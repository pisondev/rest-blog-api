package repository

import (
	"context"
	"database/sql"
	"rest-blog-api/model/domain"
)

type ArticleRepository interface {
	CreateArticle(ctx context.Context, tx *sql.Tx, article domain.Article) domain.Article
	FindArticles(ctx context.Context, tx *sql.Tx, articleFilter domain.ArticleFilter) []domain.Article
	FindById(ctx context.Context, tx *sql.Tx, articleId int) (domain.Article, error)
	UpdateById(ctx context.Context, tx *sql.Tx, article domain.Article) domain.Article
	DeleteById(ctx context.Context, tx *sql.Tx, articleId int)
}
