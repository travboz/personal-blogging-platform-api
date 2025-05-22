package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/travboz/backend-projects/personal-blog-api/internal/data/models"
	"github.com/travboz/backend-projects/personal-blog-api/internal/store/repository"
)

func createArticleHandler(repo repository.Store, logger *slog.Logger) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Content string   `json:"content"`
			Tags    []string `json:"tags"`
		}

		err := readJSON(w, r, &input)
		if err != nil {
			badRequestResponse(logger, w, r, err)
			return
		}

		article := &models.Article{
			Content:   input.Content,
			CreatedAt: time.Now(),
			Tags:      input.Tags,
		}

		err = repo.Insert(r.Context(), article)
		if err != nil {
			serverErrorResponse(logger, w, r, err)
		}

		err = writeJSON(w, http.StatusOK, envelope{"article added successfully": article}, nil)
		if err != nil {
			serverErrorResponse(logger, w, r, err)
		}
	})
}

func fetchAllArticlesHandler(repo repository.Store, logger *slog.Logger) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		articles, err := repo.FetchAllArticles(r.Context())
		if err != nil {
			serverErrorResponse(logger, w, r, err)
		}

		err = writeJSON(w, http.StatusOK, envelope{"articles": articles}, nil)
		if err != nil {
			serverErrorResponse(logger, w, r, err)
		}
	})
}
