package repository

import (
	"context"
	"database/sql"
	"rest-blog-api/helper"
	"rest-blog-api/model/domain"
)

type ArticleRepositoryImpl struct {
}

func NewArticleRepository() ArticleRepository {
	return &ArticleRepositoryImpl{}
}

func (repository *ArticleRepositoryImpl) CreateArticle(ctx context.Context, tx *sql.Tx, article domain.Article) domain.Article {
	SQL := "INSERT INTO articles(title, content) VALUES (?,?)"
	result, err := tx.ExecContext(ctx, SQL, article.Title, article.Content)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	article.Id = int(id)

	return article
}
