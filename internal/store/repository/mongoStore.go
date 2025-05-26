package repository

import (
	"context"
	"errors"
	"fmt"
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

func NewMongoStore(dbName string, db *mongo.Client) (*MongoStore, error) {
	col := db.Database(dbName).Collection("articles")

	// ensure indexes
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mod := mongo.IndexModel{
		Keys: bson.D{{Key: "content", Value: "text"}}, // example index
	}
	_, err := col.Indexes().CreateOne(ctx, mod)
	if err != nil {
		return nil, fmt.Errorf("failed to create index: %w", err)
	}

	return &MongoStore{
		articles: col,
	}, nil
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

func (m *MongoStore) FetchAllArticles(ctx context.Context, content string, tags []string) ([]*data.Article, error) {

	// filtering using content, works with single terms and phrases
	filter := bson.M{}

	if content != "" {
		filter["content"] = bson.M{
			"$regex": primitive.Regex{
				Pattern: content,
				Options: "i",
			},
		}
	}

	if len(tags) > 0 {
		filter["tags"] = bson.M{
			"$all": tags,
		}
	}

	cursor, err := m.articles.Find(ctx, filter)
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
