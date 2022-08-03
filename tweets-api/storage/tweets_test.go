package storage

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"tweets/storage/mongo/aggregations"
)

type mongoDbClientMock struct {
	mock.Mock
}

func (mc *mongoDbClientMock) DropCollection() error {
	args := mc.Called()
	return args.Error(1)
}

func (mc *mongoDbClientMock) InitCollection(data []interface{}) error {
	args := mc.Called(data)
	return args.Error(1)
}

func (mc *mongoDbClientMock) Aggregate(pipelineBuilder aggregations.IPipelinesBuilder) ([]interface{}, error) {
	args := mc.Called(pipelineBuilder)
	return args.Get(0).([]interface{}), args.Error(1)
}

func Test_GetAggregatedResult_ReturnResult(t *testing.T) {
	// Arrange
	dbClientMock := new(mongoDbClientMock)
	groupBy := []string{"year"}
	pipelineBuilder := aggregations.NewBuilder(groupBy)

	aggregation1 := TweetAggregation{
		TotalReplies:    1,
		TweetsCount:     2,
		TotalLikes:      10,
		AverageReplies:  1,
		AverageLikes:    1,
		AverageRetweets: 2,
		Label:           "2020-year",
	}

	tweetsAggregations := make([]interface{}, 0)
	tweetsAggregations = append(tweetsAggregations, aggregation1)
	dbClientMock.On("Aggregate", pipelineBuilder).Return(tweetsAggregations, nil)
	tweets := Tweet{DbClient: dbClientMock}

	// Act
	result, err := tweets.GetAggregatedResult(pipelineBuilder)

	//Assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, aggregation1.TweetsCount, result[0].TweetsCount)
	assert.Equal(t, aggregation1.Label, result[0].Label)
	assert.Equal(t, aggregation1.AverageRetweets, result[0].AverageRetweets)
	assert.Equal(t, aggregation1.AverageRetweets, result[0].AverageRetweets)
	assert.Equal(t, aggregation1.AverageReplies, result[0].AverageReplies)
	assert.Equal(t, aggregation1.AverageLikes, result[0].AverageLikes)
	assert.Equal(t, aggregation1.TotalReplies, result[0].TotalReplies)
	assert.Equal(t, aggregation1.TotalLikes, result[0].TotalLikes)
	assert.Equal(t, aggregation1.TotalRetweets, result[0].TotalRetweets)
}

func Test_GetAggregatedResult_EmptyResult(t *testing.T) {
	// Arrange
	dbClientMock := new(mongoDbClientMock)
	groupBy := []string{"year"}
	pipelineBuilder := aggregations.NewBuilder(groupBy)

	tweetsAggregations := make([]interface{}, 0)
	dbClientMock.On("Aggregate", pipelineBuilder).Return(tweetsAggregations, nil)
	tweets := Tweet{DbClient: dbClientMock}

	// Act
	result, err := tweets.GetAggregatedResult(pipelineBuilder)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func Test_GetAggregatedResult_Error(t *testing.T) {
	// Arrange
	dbClientMock := new(mongoDbClientMock)
	groupBy := []string{"year"}
	pipelineBuilder := aggregations.NewBuilder(groupBy)

	tweetsAggregations := make([]interface{}, 0)
	dbClientMock.On("Aggregate", pipelineBuilder).Return(tweetsAggregations, fmt.Errorf("error"))
	tweets := Tweet{DbClient: dbClientMock}

	// Act
	result, err := tweets.GetAggregatedResult(pipelineBuilder)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, result)
}
