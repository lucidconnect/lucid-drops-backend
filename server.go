package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"inverse.so/addresswatcher"
	"inverse.so/dbutils"
	"inverse.so/engine/whitelist"
	"inverse.so/graph"
	"inverse.so/internal"
	"inverse.so/route"
	"inverse.so/services"
	"inverse.so/utils"
)

const (
	defaultPort = "8080"
)

func main() {
	os.Setenv("APP_ENV", "development")
	utils.SetUpDefaultLogger()
	utils.LoadEnvironmentVariables()
	utils.SetUpLoggerFromConfig()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dsn, present := os.LookupEnv("DATABASE_URL")
	if !present {
		log.Fatal().Msg("DATABASE_URL not set")
	}

	dbutils.SetupDB(dsn)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Options{})
	srv.Use(extension.Introspection{})

	router := chi.NewRouter()
	loadCORS(router)

	router.Use(internal.UserAuthMiddleWare())
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/health", http.HandlerFunc(route.HealthCheckHandler))
	router.Handle("/query", srv)
	router.Handle("/metadata/{contract}/{itemid}", http.HandlerFunc(route.MetadataHandler))
	router.Handle("/twitter01/callback/", http.HandlerFunc(route.TwitterCallBack))
	router.Handle("/telegram/callback/", http.HandlerFunc(route.TelegramCallBack))
	router.Handle("/patreon/callback/", http.HandlerFunc(route.PatreonCallBack))
	router.Handle("/whitelist/patreon/callback/", http.HandlerFunc(route.PatreonWhitelistCallBack))

	log.Info().Msgf("connect to http://localhost:%s/ for GraphQL playground", port)

	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go createTelegramBotInstance()
	go addresswatcher.SubscribeToInverseContractDeployments()
	log.Err(httpServer.ListenAndServe())
}

func loadCORS(router *chi.Mux) {
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*", "ws://*", "wss://*", "*"},
		AllowedMethods: []string{
			http.MethodOptions,
			http.MethodGet,
			http.MethodPost,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	}).Handler)
}

func createTelegramBotInstance() {
	whitelist.InverseBot = services.InitTelegramBot()
}
