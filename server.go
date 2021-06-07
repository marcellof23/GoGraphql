package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/marcellof23/GoGraphql/configs/database"
	"github.com/marcellof23/GoGraphql/graph"
	"github.com/marcellof23/GoGraphql/graph/generated"
	"github.com/marcellof23/GoGraphql/internal/models"
)

const defaultPort = "8080"

func main() {
	err := database.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	database.DB.AutoMigrate(&models.User{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:8082/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
