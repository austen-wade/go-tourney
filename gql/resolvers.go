package gql

import (
	"github.com/austen-wade/go-graphql/postgres"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	db *postgres.Db
}

func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if ok {
		users := r.db.GetUsersByID(id)
		return users, nil
	}

	name, ok := p.Args["name"].(string)
	if ok {
		users := r.db.GetUsersByName(name)
		return users, nil
	}

	return nil, nil
}
