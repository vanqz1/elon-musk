package storage

// TweetAggregation struct contains all tweet data after aggregation
type TweetAggregation struct {
	Label           string  `bson:"label" json:"label"`
	TotalLikes      int     `bson:"totalLikes" json:"totalLikes"`
	AverageLikes    float64 `bson:"averageLikes" json:"averageLikes"`
	TotalRetweets   int     `bson:"totalRetweets" json:"totalRetweets"`
	AverageRetweets float64 `bson:"averageRetweets" json:"averageRetweets"`
	TotalReplies    int     `bson:"totalReplies" json:"totalReplies"`
	AverageReplies  float64 `bson:"averageReplies" json:"averageReplies"`
	TweetsCount     int     `bson:"tweetsCount" json:"tweetsCount""`
}
