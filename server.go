package main

import (
	"github.com/nicotanzil/backend-gqlgen/app/middleware"
	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/resolver"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nicotanzil/backend-gqlgen/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	wrapped := middleware.CorsMiddleware(srv)
	wrapped = middleware.AuthMiddleware(wrapped)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", wrapped)

	database.Migrate()
	database.Seed()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
