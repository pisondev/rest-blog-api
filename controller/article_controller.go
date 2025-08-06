package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ArticleController interface {
	CreateArticle(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAllArticles(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	DeleteById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
