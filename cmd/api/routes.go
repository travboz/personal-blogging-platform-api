package main

import (
	"log/slog"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/travboz/backend-projects/personal-blog-api/internal/store/repository"
)

func routes(logger *slog.Logger, store repository.Store) http.Handler {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/health", healthcheckHandler(logger))
	router.Handler(http.MethodPost, "/articles", createArticleHandler(logger, store))
	router.Handler(http.MethodGet, "/articles", fetchAllArticlesHandler(logger, store))
	router.Handler(http.MethodGet, "/articles/:id", fetchAllArticlesHandler(logger, store))

	return router
}
