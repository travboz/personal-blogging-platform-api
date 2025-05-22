package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/travboz/backend-projects/personal-blog-api/internal/db"
	"github.com/travboz/backend-projects/personal-blog-api/internal/env"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/health", healthcheck)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mongo_uri := env.GetString(
		"MONGODB_URI",
		"mongodb://travis:secret@localhost:27002/blog_articles?authSource=admin&readPreference=primary&appname=MongDB%20Compass&directConnection=true&ssl=false",
	)

	_, err := db.NewMongoDBClient(mongo_uri)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("mongodb successfully connected")

	port := ":7666"

	logger.Info("running server", "port", port)
	if err := http.ListenAndServe(port, router); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
