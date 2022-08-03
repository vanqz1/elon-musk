package services

type TweetAggregation struct {
	Label           string  `json:"label"`
	TotalLikes      int     `json:"totalLikes"`
	AverageLikes    float64 `json:"averageLikes"`
	TotalRetweets   int     `json:"totalRetweets"`
	AverageRetweets float64 `json:"averageRetweets"`
	TotalReplies    int     `json:"totalReplies"`
	AverageReplies  float64 `json:"averageReplies"`
	TweetsCount     int     `json:"tweetsCount""`
}
