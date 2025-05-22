package repository

import (
	"context"
	"time"

	"github.com/travboz/backend-projects/personal-blog-api/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStore struct {
	articles *mongo.Collection
}

func NewMongoStore(dbName string, db *mongo.Client) *MongoStore {
	col := db.Database(dbName).Collection("articles")

	return &MongoStore{
		articles: col,
	}
}

func (m *MongoStore) Insert(ctx context.Context, article *models.Article) error {
	article.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := m.articles.InsertOne(ctx, article)

	return err
}

func (m *MongoStore) GetArticleById(context.Context, string) (*models.Article, error) {
	return nil, nil
}

func (m *MongoStore) FetchAllArticles(ctx context.Context) ([]*models.Article, error) {
	cursor, err := m.articles.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var articles []*models.Article

	for cursor.Next(ctx) {
		var article *models.Article
		if err := cursor.Decode(&article); err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func (m *MongoStore) UpdateArtcle(context.Context, string, models.Article) error {
	return nil
}

func (m *MongoStore) DeleteArticle(context.Context, string) error {
	return nil
}

func (m *MongoStore) Shutdown(context.Context) error {
	return nil
}
