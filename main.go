package main

import (
	"net/http"
	"rest-blog-api/app"
	"rest-blog-api/controller"
	"rest-blog-api/exception"
	"rest-blog-api/helper"
	"rest-blog-api/repository"
	"rest-blog-api/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	articleRepository := repository.NewArticleRepository()
	articleService := service.NewArticleService(articleRepository, db, validate)
	articleController := controller.NewArticleController(articleService)

	router := httprouter.New()

	router.POST("/api/articles", articleController.CreateArticle)
	router.GET("/api/articles", articleController.FindAllArticles)
	router.GET("/api/articles/:articleId", articleController.FindById)
	router.PUT("/api/articles/:articleId", articleController.UpdateById)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
