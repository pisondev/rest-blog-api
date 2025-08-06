package service

import (
	"context"
	"database/sql"
	"rest-blog-api/helper"
	"rest-blog-api/model/domain"
	"rest-blog-api/model/web"
	"rest-blog-api/repository"

	"github.com/go-playground/validator/v10"
)

type ArticleServiceImpl struct {
	ArticleRepository repository.ArticleRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewArticleService(articleRepository repository.ArticleRepository, DB *sql.DB, validate *validator.Validate) ArticleService {
	return &ArticleServiceImpl{
		ArticleRepository: articleRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ArticleServiceImpl) CreateArticle(ctx context.Context, req web.ArticleCreateRequest) web.ArticleResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	article := domain.Article{
		Title:   req.Title,
		Content: req.Content,
	}
	createdArticle := service.ArticleRepository.CreateArticle(ctx, tx, article)

	return helper.ToArticleResponse(createdArticle)
}

func (service *ArticleServiceImpl) FindAllArticles(ctx context.Context) []web.ArticleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	articles := service.ArticleRepository.FindAllArticles(ctx, tx)

	return helper.ToArticleResponses(articles)
}
