package exception

import (
	"net/http"
	"rest-blog-api/helper"
	"rest-blog-api/model/web"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	internalServerError(w, r, err)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(w, webResponse)
}
