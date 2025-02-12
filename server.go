package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	database "posts/db"
	"posts/graph"
	"posts/inmemory"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8081"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	storageType := os.Getenv("STORAGE")
	if storageType == "" {
		storageType = "postgresql"
	}

	var resolver *graph.Resolver
	fmt.Println(storageType)
	if storageType == "postgresql" {
		database.ConnectDB()
		database.CreateDB()
		database.MigrateDB()
		resolver = &graph.Resolver{Database: database.DBInstance}
	} else {
		resolver = &graph.Resolver{InMemoryStorage: inmemory.NewInMemoryDB()}
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
