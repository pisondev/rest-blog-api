package helper

import (
	"rest-blog-api/model/domain"
	"rest-blog-api/model/web"
)

func ToArticleResponse(article domain.Article) web.ArticleResponse {
	return web.ArticleResponse{
		Id:      article.Id,
		Title:   article.Title,
		Content: article.Content,
	}
}

func ToArticleResponses(articles []domain.Article) []web.ArticleResponse {
	var articleResponses []web.ArticleResponse
	for _, article := range articles {
		articleResponses = append(articleResponses, ToArticleResponse(article))
	}
	return articleResponses
}
