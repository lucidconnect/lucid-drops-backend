package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"inverse.so/graph"
	"inverse.so/internal"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	router := chi.NewRouter()
	router.Use(internal.UserAuthMiddleWare())
	loadCORS(router)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func loadCORS(router *chi.Mux) {
	switch os.Getenv("APP_ENV") {
	// TODO - add proper CORS support for production & staging
	default:
		router.Use(cors.New(cors.Options{
			AllowedOrigins: []string{"https://*", "http://*", "ws://*", "wss://*"},
			AllowedMethods: []string{
				http.MethodOptions,
				http.MethodGet,
				http.MethodPost,
			},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: false,
		}).Handler)
	}
}
