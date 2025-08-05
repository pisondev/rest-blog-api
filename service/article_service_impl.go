package service

import (
	"context"
	"database/sql"
	"rest-blog-api/helper"
	"rest-blog-api/model/domain"
	"rest-blog-api/model/web"
	"rest-blog-api/repository"
)

type ArticleServiceImpl struct {
	ArticleRepository repository.ArticleRepository
	DB                *sql.DB
}

func (service *ArticleServiceImpl) CreateArticle(ctx context.Context, req web.ArticleCreateRequest) web.ArticleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	helper.CommitOrRollback(tx)

	article := domain.Article{
		Title:   req.Title,
		Content: req.Content,
	}
	createdArticle := service.ArticleRepository.CreateArticle(ctx, tx, article)

	return helper.ToArticleResponse(createdArticle)
}
