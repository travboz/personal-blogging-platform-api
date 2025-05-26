package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/travboz/backend-projects/personal-blog-api/internal/data"
	"github.com/travboz/backend-projects/personal-blog-api/internal/store/repository"
	"github.com/travboz/backend-projects/personal-blog-api/internal/validator"
)

func createArticleHandler(logger *slog.Logger, store repository.Store) http.Handler {

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

		article := &data.Article{
			Content:   input.Content,
			CreatedAt: time.Now(),
			Tags:      input.Tags,
		}

		v := validator.New()

		if data.ValidateArticle(v, article); !v.Valid() {
			failedValidationResponse(logger, w, r, v.Errors)
			return
		}

		err = store.Insert(r.Context(), article)
		if err != nil {
			serverErrorResponse(logger, w, r, err)
		}

		// telling the client where to find the new article at:
		headers := make(http.Header)
		headers.Set("Location", fmt.Sprintf("/v1/movies/%s", article.ID.Hex()))

		err = writeJSON(w, http.StatusOK, envelope{"article": article}, headers)
		if err != nil {
			serverErrorResponse(logger, w, r, err)
		}
	})
}

func fetchAllArticlesHandler(logger *slog.Logger, store repository.Store) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Content string
			Tags    []string
		}

		qs := r.URL.Query()

		input.Content = readString(qs, "content", "")
		input.Tags = readTags(qs, "tags", []string{})

		articles, err := store.FetchAllArticles(r.Context(), input.Content, input.Tags)
		if err != nil {
			serverErrorResponse(logger, w, r, err)
		}

		err = writeJSON(w, http.StatusOK, envelope{"articles": articles}, nil)
		if err != nil {
			serverErrorResponse(logger, w, r, err)
		}
	})
}

func getArticleByIdHandler(logger *slog.Logger, store repository.Store) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := readIDParam(r)

		article, err := store.GetArticleById(r.Context(), id)
		if err != nil {
			switch {
			case errors.Is(err, repository.ErrRecordNotFound):
				notFoundResponse(logger, w, r)
			default:
				serverErrorResponse(logger, w, r, err)
			}

			return
		}

		err = writeJSON(w, http.StatusOK, envelope{"article": article}, nil)
		if err != nil {
			serverErrorResponse(logger, w, r, err)
		}

	})
}

func deleteArticleHandler(logger *slog.Logger, store repository.Store) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := readIDParam(r)

		err := store.DeleteArticle(r.Context(), id)
		if err != nil {
			switch {
			case errors.Is(err, repository.ErrRecordNotFound):
				notFoundResponse(logger, w, r)
			default:
				serverErrorResponse(logger, w, r, err)
			}

			return
		}

		err = writeJSON(w, http.StatusOK, envelope{"message": "succesful deletion of article with id: " + id}, nil)
		if err != nil {
			serverErrorResponse(logger, w, r, err)
		}

	})
}

func updateArticleHandler(logger *slog.Logger, store repository.Store) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := readIDParam(r)

		var input struct {
			Content string   `json:"content" bson:"content"`
			Tags    []string `json:"tags" bson:"tags"`
		}

		err := readJSON(w, r, &input)
		if err != nil {
			badRequestResponse(logger, w, r, err)
			return
		}

		article := &data.Article{
			Content: input.Content,
			Tags:    input.Tags,
		}

		v := validator.New()

		if data.ValidateArticle(v, article); !v.Valid() {
			failedValidationResponse(logger, w, r, v.Errors)
			return
		}

		updated, err := store.UpdateArtcle(r.Context(), id, article)
		if err != nil {
			switch {
			case errors.Is(err, repository.ErrRecordNotFound):
				notFoundResponse(logger, w, r)
			default:
				serverErrorResponse(logger, w, r, err)
			}

			return
		}

		err = writeJSON(w, http.StatusOK, envelope{"article": updated}, nil)
		if err != nil {
			serverErrorResponse(logger, w, r, err)
		}

	})
}
