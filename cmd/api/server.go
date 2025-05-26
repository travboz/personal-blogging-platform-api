package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/travboz/backend-projects/personal-blog-api/internal/env"
)

func serve(logger *slog.Logger, router http.Handler) error {

	srv := &http.Server{
		Addr:         env.GetString("SERVER_PORT", ":7666"),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr)

	return srv.ListenAndServe()
}
