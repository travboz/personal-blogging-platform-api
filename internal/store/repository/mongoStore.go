package repository

import (
	"context"
	"errors"
	"time"

	"github.com/travboz/backend-projects/personal-blog-api/internal/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (m *MongoStore) Insert(ctx context.Context, article *data.Article) error {
	article.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := m.articles.InsertOne(ctx, article)

	return err
}

func (m *MongoStore) GetArticleById(ctx context.Context, id string) (*data.Article, error) {

	article_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := m.articles.FindOne(ctx, bson.M{"_id": article_id})

	var article data.Article
	if err = result.Decode(&article); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &article, nil
}

func (m *MongoStore) FetchAllArticles(ctx context.Context) ([]*data.Article, error) {
	cursor, err := m.articles.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var articles []*data.Article

	for cursor.Next(ctx) {
		var article *data.Article
		if err := cursor.Decode(&article); err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func (m *MongoStore) UpdateArtcle(ctx context.Context, id string, article *data.Article) (*data.Article, error) {
	article_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err // invalid ID format
	}

	filter := bson.M{"_id": article_id}
	update := bson.D{
		{"$set", bson.D{
			{"content", article.Content},
			{"tags", article.Tags},
		}},
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := m.articles.FindOneAndUpdate(
		ctx,
		filter,
		update,
		&opt,
	)

	var updatedArticle data.Article

	if err = result.Decode(&updatedArticle); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &updatedArticle, nil
}

func (m *MongoStore) DeleteArticle(ctx context.Context, id string) error {
	article_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := m.articles.DeleteOne(ctx, bson.M{"_id": article_id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m *MongoStore) Shutdown(context.Context) error {
	return nil
}
