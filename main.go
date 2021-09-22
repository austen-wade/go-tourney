package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/austen-wade/go-graphql/gql"
	"github.com/austen-wade/go-graphql/postgres"
	"github.com/austen-wade/go-graphql/server"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "tourneys"
)

func main() {
	router, db := initializeAPI()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":4000", router))
}

func initializeAPI() (*chi.Mux, *postgres.Db) {
	router := chi.NewRouter()

	db, err := postgres.New(
		postgres.ConnString(host, port, user, password, dbname),
	)
	if err != nil {
		log.Fatal(err)
	}

	rootQuery := gql.NewRoot(db)
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	s := server.Server{
		GqlSchema: &sc,
	}

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.Compress(5, "gzip"),
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	router.Post("/graphql", s.GraphQl())

	return router, db
}
