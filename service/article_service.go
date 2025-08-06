package service

import (
	"context"
	"rest-blog-api/model/web"
)

type ArticleService interface {
	CreateArticle(ctx context.Context, req web.ArticleCreateRequest) web.ArticleResponse
	FindAllArticles(ctx context.Context) []web.ArticleResponse
}
