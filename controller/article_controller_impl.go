package controller

import (
	"net/http"
	"rest-blog-api/helper"
	"rest-blog-api/model/web"
	"rest-blog-api/service"
	"strconv"

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
	helper.ReadFromRequestBody(r, &result)

	articleResponse := controller.ArticleService.CreateArticle(r.Context(), result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	webResponse := web.WebResponse{
		Code:   201,
		Status: "Created",
		Data:   articleResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *ArticleControllerImpl) FindAllArticles(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	articleResponses := controller.ArticleService.FindAllArticles(r.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   articleResponses,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *ArticleControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	articleId := params.ByName("articleId")
	id, err := strconv.Atoi(articleId)
	helper.PanicIfError(err)

	articleResponse := controller.ArticleService.FindById(r.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   articleResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}
