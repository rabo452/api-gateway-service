package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	APP_PORT := os.Getenv("APP_PORT")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	http.ListenAndServe(":"+APP_PORT, r)
}
