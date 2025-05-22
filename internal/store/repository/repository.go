package repository

import (
	"context"

	"github.com/travboz/backend-projects/personal-blog-api/internal/data/models"
)

type Store interface {
	Insert(context.Context, *models.Article) error
	GetArticleById(context.Context, string) (*models.Article, error)
	FetchAllArticles(context.Context) ([]*models.Article, error)
	UpdateArtcle(context.Context, string, models.Article) error
	DeleteArticle(context.Context, string) error
	Shutdown(context.Context) error
}
