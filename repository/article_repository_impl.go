package repository

import (
	"context"
	"database/sql"
	"rest-blog-api/model/domain"
)

type ArticleRepositoryImpl struct {
}

func (repository *ArticleRepositoryImpl) CreateArticle(ctx context.Context, tx *sql.Tx, article domain.Article) domain.Article {
	SQL := "INSERT INTO articles(title, content) VALUES (?,?)"
	result, err := tx.ExecContext(ctx, SQL, article.Title, article.Content)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	article.Id = int(id)

	return article
}
