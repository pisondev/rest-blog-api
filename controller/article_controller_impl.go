package controller

import (
	"net/http"
	"rest-blog-api/helper"
	"rest-blog-api/model/web"
	"rest-blog-api/service"

	"github.com/julienschmidt/httprouter"
)

type ArticleControllerImpl struct {
	ArticleService service.ArticleService
}

func NewArticleController(articleService service.ArticleService) ArticleController {
	return &ArticleControllerImpl{
		ArticleService: articleService,
	}
}

func (controller *ArticleControllerImpl) CreateArticle(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	result := web.ArticleCreateRequest{}
	helper.ReadFromRequestBody(r, result)

	articleResponse := controller.ArticleService.CreateArticle(r.Context(), result)

	webResponse := web.WebResponse{
		Code:   201,
		Status: "Created",
		Data:   articleResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}
