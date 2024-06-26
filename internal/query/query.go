package query

import (
	"cqrs-layout-example/internal/query/entity_name"
	"cqrs-layout-example/internal/query/entity_name/fetch"
	"cqrs-layout-example/internal/query/entity_name/get_by_id"
)

type EntityQueries struct {
	GetEventByIDQuery    *get_by_id.QueryHandler
	FetchEventsByIDQuery *fetch.QueryHandler
}

func NewOrderQueries(storage entity_name.StorageInterface) *EntityQueries {
	return &EntityQueries{
		GetEventByIDQuery:    get_by_id.NewQueryHandler(storage),
		FetchEventsByIDQuery: fetch.NewQueryHandler(storage),
	}
}
