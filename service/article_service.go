package service

import (
	"context"
	"rest-blog-api/model/web"
)

type ArticleService interface {
	CreateArticle(ctx context.Context, req web.ArticleCreateRequest) web.ArticleResponse
	FindAllArticles(ctx context.Context) []web.ArticleResponse
	FindById(ctx context.Context, categoryId int) web.ArticleResponse
	UpdateById(ctx context.Context, req web.ArticleUpdateRequest) web.ArticleResponse
}
