package storage

import (
	"encoding/json"
	"tweets/storage/mongo/aggregations"
)

type ITweet interface {
	GetAggregatedResult(pipelineBuilder aggregations.IPipelinesBuilder) ([]TweetAggregation, error)
}

// Tweet wraps all queries for tweets
type Tweet struct {
	DbClient IDbClient
}

// NewTweet creates new instance of Tweet
func NewTweet(client IDbClient) ITweet {
	return &Tweet{
		DbClient: client,
	}
}

// GetAggregatedResult returns aggregation result
func (tm *Tweet) GetAggregatedResult(pipelineBuilder aggregations.IPipelinesBuilder) ([]TweetAggregation, error) {
	result := make([]TweetAggregation, 0)
	data, err := tm.DbClient.Aggregate(pipelineBuilder)
	if err != nil {
		return nil, err
	}

	for _, record := range data {
		// there is an issue casting bson, marshal unmarshal is workaround
		a, err := json.Marshal(record)
		var res TweetAggregation
		if err != nil {
			return result, err
		}

		json.Unmarshal(a, &res)
		result = append(result, res)
	}

	return result, nil
}
