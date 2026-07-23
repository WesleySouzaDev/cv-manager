package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/WesleySouzaDev/cv-manager/config"
	"github.com/WesleySouzaDev/cv-manager/generated"
	"github.com/WesleySouzaDev/cv-manager/resolvers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	env := config.GetEnv()
	port := env.Port
	if port == "" {
		port = config.DEFAULT_PORT
	}

	router := chi.NewRouter()
	router.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins:   []string{"*"},
				AllowedMethods:   []string{http.MethodPost, http.MethodGet, http.MethodOptions},
				AllowCredentials: true,
			},
		),
	)

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
