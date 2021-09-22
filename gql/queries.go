package gql

import (
	"github.com/austen-wade/go-graphql/postgres"
	"github.com/graphql-go/graphql"
)

type Root struct {
	Query    *graphql.Object
	Mutation *graphql.Object
}

func NewRoot(db *postgres.Db) *Root {
	resolver := Resolver{db: db}

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"events": &graphql.Field{
						Type: graphql.NewList(Event),
						Args: graphql.FieldConfigArgument{
							"event_id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.EventResolver,
					},
					"entrants": &graphql.Field{
						Type: graphql.NewList(Entrant),
						Args: graphql.FieldConfigArgument{
							"entrant_id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.EntrantResolver,
					},
					"sets": &graphql.Field{
						Type: graphql.NewList(Set),
						Args: graphql.FieldConfigArgument{
							"set_id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.SetResolver,
					},
				},
			},
		),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"createSet": &graphql.Field{
					Type: Set,
					Args: graphql.FieldConfigArgument{
						"entrant1_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"entrant2_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"entrant1_result": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"entrant2_result": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: resolver.CreateSetResolver,
				},
			},
		}),
	}
	return &root
}
