package gql

import (
	"github.com/austen-wade/go-graphql/postgres"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	db *postgres.Db
}

func (r *Resolver) EventResolver(p graphql.ResolveParams) (interface{}, error) {
	eventID, ok := p.Args["event_id"].(int)
	if ok {
		events := r.db.GetEventsByID(eventID)
		return events, nil
	}

	return nil, nil
}

func (r *Resolver) EntrantResolver(p graphql.ResolveParams) (interface{}, error) {
	entrantID, ok := p.Args["entrant_id"].(int)
	if ok {
		entrants := r.db.GetEntrantsByID(entrantID)
		return entrants, nil
	}

	return nil, nil
}

func (r *Resolver) SetResolver(p graphql.ResolveParams) (interface{}, error) {
	setID, ok := p.Args["set_id"].(int)
	if ok {
		sets := r.db.GetSetsByID(setID)
		return sets, nil
	}

	return nil, nil
}

/* func (r *Resolver) CreateEventResolver(p graphql.ResolveParams) (interface{}, error) {

}
func (r *Resolver) CreateEntrantResolver(p graphql.ResolveParams) (interface{}, error) {

} */

func (r *Resolver) CreateSetResolver(p graphql.ResolveParams) (interface{}, error) {
	setParams, ok := p.Source.(postgres.SetParams)
	if ok {
		id := r.db.SaveSet(setParams)
		return id, nil
	}

	return nil, nil
}
