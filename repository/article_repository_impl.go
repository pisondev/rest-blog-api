package repository

import (
	"context"
	"database/sql"
	"errors"
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

func (repository *ArticleRepositoryImpl) FindArticles(ctx context.Context, tx *sql.Tx, articleFilter domain.ArticleFilter) []domain.Article {
	SQL := "SELECT id, title, content, created_at, updated_at FROM articles WHERE 1=1"
	var args []interface{}

	if articleFilter.Title != "" {
		SQL += " AND title LIKE ?"
		args = append(args, "%"+articleFilter.Title+"%")
	}
	if !articleFilter.StartDate.IsZero() {
		SQL += " AND created_at >= ?"
		args = append(args, articleFilter.StartDate)
	}
	if !articleFilter.EndDate.IsZero() {
		SQL += " AND created_at <= ?"
		args = append(args, articleFilter.EndDate)
	}

	rows, err := tx.QueryContext(ctx, SQL, args...)
	helper.PanicIfError(err)
	defer rows.Close()

	var articles []domain.Article
	for rows.Next() {
		article := domain.Article{}
		err := rows.Scan(&article.Id, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt)
		helper.PanicIfError(err)

		articles = append(articles, article)
	}

	return articles
}

func (repository *ArticleRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, articleId int) (domain.Article, error) {
	SQL := "SELECT id, title, content, created_at, updated_at FROM articles WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, articleId)
	helper.PanicIfError(err)
	defer rows.Close()

	article := domain.Article{}
	if rows.Next() {
		err := rows.Scan(&article.Id, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt)
		helper.PanicIfError(err)
		return article, nil
	} else {
		return article, errors.New("article is not found")
	}
}

func (repository *ArticleRepositoryImpl) UpdateById(ctx context.Context, tx *sql.Tx, article domain.Article) domain.Article {
	SQL := "UPDATE articles SET title = ?, content = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, article.Title, article.Content, article.Id)
	helper.PanicIfError(err)

	return article
}

func (repository *ArticleRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, articleId int) {
	SQL := "DELETE FROM articles WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, articleId)
	helper.PanicIfError(err)
}
