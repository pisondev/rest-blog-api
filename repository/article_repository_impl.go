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

func (repository *ArticleRepositoryImpl) FindAllArticles(ctx context.Context, tx *sql.Tx) []domain.Article {
	SQL := "SELECT id, title, content FROM articles"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var articles []domain.Article
	for rows.Next() {
		article := domain.Article{}
		err := rows.Scan(&article.Id, &article.Title, &article.Content)
		helper.PanicIfError(err)

		articles = append(articles, article)
	}

	return articles
}
