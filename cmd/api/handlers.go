package main

import (
	"fmt"
	"net/http"
)

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "service up\n")
}
