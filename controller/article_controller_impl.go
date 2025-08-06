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

func (controller *ArticleControllerImpl) FindArticles(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	articleResponses := controller.ArticleService.FindArticles(r.Context())

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

func (controller *ArticleControllerImpl) UpdateById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	articleUpdateReq := web.ArticleUpdateRequest{}
	helper.ReadFromRequestBody(r, &articleUpdateReq)

	articleId := params.ByName("articleId")
	id, err := strconv.Atoi(articleId)
	helper.PanicIfError(err)

	articleUpdateReq.Id = id

	articleResponse := controller.ArticleService.UpdateById(r.Context(), articleUpdateReq)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   articleResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *ArticleControllerImpl) DeleteById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	articleId := params.ByName("articleId")
	id, err := strconv.Atoi(articleId)
	helper.PanicIfError(err)

	controller.ArticleService.DeleteById(r.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(w, webResponse)
}
