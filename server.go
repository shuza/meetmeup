package main

import (
	"github.com/go-pg/pg/v10"
	"github.com/shuza/meetmeup/postgres"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/shuza/meetmeup/graphql"
)

const defaultPort = "8080"

func main() {
	DB := postgres.New(&pg.Options{
		User:     "foobar",
		Password: "foobar",
		Database: "meetup_dev",
	})
	defer DB.Close()
	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	conf := graphql.Config{Resolvers: &graphql.Resolver{
		MeetupsRepo: postgres.MeetupsRepo{DB: DB},
		UsersRepo:   postgres.UsersRepo{DB: DB},
	}}

	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(conf))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
