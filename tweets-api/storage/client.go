package storage

import (
	"tweets/storage/mongo/aggregations"
)

type IDbClient interface {
	DropCollection() error
	InitCollection(data []interface{}) error
	Aggregate(builder aggregations.IPipelinesBuilder) ([]interface{}, error)
}
