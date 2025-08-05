package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ArticleController interface {
	CreateArticle(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
