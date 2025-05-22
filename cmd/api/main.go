package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
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

	store := repository.NewMongoStore(env.GetString("MONGO_DB_NAME", "blog_articles"), mongoClient)

	router := httprouter.New()

	router.Handler(http.MethodGet, "/health", healthcheckHandler(logger))
	router.Handler(http.MethodPost, "/articles", createArticleHandler(store, logger))
	router.Handler(http.MethodGet, "/articles", fetchAllArticlesHandler(store, logger))

	srv := &http.Server{
		Addr:         env.GetString("SERVER_PORT", ":7666"),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("running server", "port", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		logger.Error(err.Error())
		mongoClient.Disconnect(context.Background())
		os.Exit(1)
	}
}
