package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/travboz/backend-projects/personal-blog-api/internal/db"
	"github.com/travboz/backend-projects/personal-blog-api/internal/env"
	"github.com/travboz/backend-projects/personal-blog-api/internal/store/repository"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mongo_uri := env.GetString(
		"MONGODB_URI",
		"mongodb://travis:secret@localhost:27002/blog_articles?authSource=admin&readPreference=primary&appname=MongDB%20Compass&directConnection=true&ssl=false",
	)

	mongoClient, err := db.NewMongoDBClient(mongo_uri)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("mongodb successfully connected")

	store, err := repository.NewMongoStore(env.GetString("MONGO_DB_NAME", "blog_articles"), mongoClient)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	router := routes(logger, store)

	err = serve(logger, router)

	if err != nil {
		logger.Error(err.Error())
		mongoClient.Disconnect(context.Background())
		os.Exit(1)
	}
}
