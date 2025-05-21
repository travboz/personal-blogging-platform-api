package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/health", healthcheck)

	fmt.Println("Running server on port 7666")
	if err := http.ListenAndServe(":7666", router); err != nil {
		fmt.Println("error encountered", err)
		os.Exit(1)
	}
}
