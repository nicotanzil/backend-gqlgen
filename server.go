package main

import (
	"github.com/nicotanzil/backend-gqlgen/graph/resolver"
	"github.com/nicotanzil/backend-gqlgen/middleware"
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
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", wrapped)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
