package gql

import "github.com/graphql-go/graphql"

var Event = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Event",
		Fields: graphql.Fields{
			"event_id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"game_name": &graphql.Field{
				Type: graphql.String,
			},
			"number_entrants": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var Entrant = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Entrant",
		Fields: graphql.Fields{
			"entrant_id": &graphql.Field{
				Type: graphql.Int,
			},
			"entrant_tag": &graphql.Field{
				Type: graphql.String,
			},
			"initial_seed": &graphql.Field{
				Type: graphql.Int,
			},
			"final_placement": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var Set = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Set",
		Fields: graphql.Fields{
			"set_id": &graphql.Field{
				Type: graphql.Int,
			},
			"entrant1_id": &graphql.Field{
				Type: graphql.Int,
			},
			"entrant2_id": &graphql.Field{
				Type: graphql.Int,
			},
			"entrant1_result": &graphql.Field{
				Type: graphql.String,
			},
			"entrant2_result": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
