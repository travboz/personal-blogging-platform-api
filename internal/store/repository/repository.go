package repository

import (
	"context"

	"github.com/travboz/backend-projects/personal-blog-api/internal/data"
)

type Store interface {
	Insert(context.Context, *data.Article) error
	GetArticleById(context.Context, string) (*data.Article, error)
	FetchAllArticles(context.Context, string, []string) ([]*data.Article, error)
	UpdateArtcle(context.Context, string, *data.Article) (*data.Article, error)
	DeleteArticle(context.Context, string) error
	Shutdown(context.Context) error
}
