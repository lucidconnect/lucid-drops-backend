package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"inverse.so/graph"
	"inverse.so/internal"
	"inverse.so/utils"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	utils.SetUpDefaultLogger()
	utils.LoadEnvironmentVariables()
	utils.SetUpLoggerFromConfig()

	dsn, present := os.LookupEnv("DATABASE_URL")
	if !present {
		log.Fatal().Msg("DATABASE_URL not set")
	}

	utils.SetupDB(dsn)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	router := chi.NewRouter()
	router.Use(internal.UserAuthMiddleWare())
	loadCORS(router)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Err(http.ListenAndServe(":"+port, nil))
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
