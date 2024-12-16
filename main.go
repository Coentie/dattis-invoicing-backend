package main

import (
	"dattis/database"
	"dattis/router"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

// main entrypoint of the application.
func main() {
	r, err := bootstrap()

	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	http.ListenAndServe(":"+os.Getenv("APP_PORT"), r)
}

// bootstrap bootstraps the applications.
func bootstrap() (*chi.Mux, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error_connecting_database.")
		return nil, err
	}

	err = database.New()

	if err != nil {
		log.Fatal("error_connecting_database")
		return nil, err
	}

	err = database.Migrate()

	if err != nil {
		log.Fatal("error_migrating_database")
		return nil, err
	}

	if os.Getenv("APP_ENV") == "local" || os.Getenv("APP_ENV") == "dev" {
		database.Seed()
	}

	return router.New(), nil
}
