package service

import (
	"context"
	"database/sql"
	"rest-blog-api/exception"
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

func (service *ArticleServiceImpl) FindById(ctx context.Context, categoryId int) web.ArticleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	article, err := service.ArticleRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToArticleResponse(article)
}

func (service *ArticleServiceImpl) UpdateById(ctx context.Context, req web.ArticleUpdateRequest) web.ArticleResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	articleId := req.Id

	article, err := service.ArticleRepository.FindById(ctx, tx, articleId)
	helper.PanicIfError(err)

	article.Title = req.Title
	article.Content = req.Content

	updatedArticle := service.ArticleRepository.UpdateById(ctx, tx, article)

	return helper.ToArticleResponse(updatedArticle)
}

func (service *ArticleServiceImpl) DeleteById(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	article, err := service.ArticleRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ArticleRepository.DeleteById(ctx, tx, article)
}
