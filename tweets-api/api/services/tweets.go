package services

import (
	"tweets/api/services/enums"
	"tweets/storage"
	"tweets/storage/mongo"
	"tweets/storage/mongo/aggregations"
)

type TweetsService struct {
	dbClient storage.IDbClient
	tweets   storage.ITweet
}

// NewTweetsService creates instance of TweetsService
func NewTweetsService() (*TweetsService, error) {
	dbConfig := mongo.NewDbConfig(TweetsDBName, TweetCollectionName)
	mongoClient, err := mongo.NewDbMongoClient(dbConfig)
	if err != nil {
		return nil, err
	}

	return &TweetsService{
		dbClient: mongoClient,
		tweets:   storage.NewTweet(mongoClient),
	}, err
}

// GetAggregatedResult returns aggregation result grouped by input
func (ts *TweetsService) GetAggregatedResult(groupBy []enums.GroupByEnum) ([]TweetAggregation, error) {
	result := make([]TweetAggregation, 0)
	groupByAsString, err := enums.ConvertToStringsGroupBy(groupBy)
	if err != nil {
		return nil, err
	}

	pipelineBuilder := aggregations.NewBuilder(groupByAsString)

	items, err := ts.tweets.GetAggregatedResult(pipelineBuilder)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		result = append(result, TweetAggregation{
			Label:           item.Label,
			AverageReplies:  item.AverageReplies,
			AverageLikes:    item.AverageLikes,
			AverageRetweets: item.AverageRetweets,
			TotalLikes:      item.TotalLikes,
			TotalReplies:    item.TotalReplies,
			TotalRetweets:   item.TotalRetweets,
			TweetsCount:     item.TweetsCount,
		})
	}

	return result, nil
}
