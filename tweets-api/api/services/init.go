package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tweets/storage/mongo"
)

const TweetsCollectionJsonUrl = "https://drive.google.com/uc?export=download&id=1IdVMuwdvSJNof8QOI7f7RJf9wFO0cpgr"
const TweetsDBName = "TweetsDb"
const TweetCollectionName = "Tweets"

// InitDatabase init database
func InitDatabase() error {
	log.Println("[Info] Database initialization started")
	dbConfig := mongo.NewDbConfig(TweetsDBName, TweetCollectionName)
	client, err := mongo.NewDbMongoClient(dbConfig)
	if err != nil {
		return fmt.Errorf("failed to get mongo client, %s", err)
	}

	// drop collection if exists
	err = client.DropCollection()
	if err != nil {
		return fmt.Errorf("failed to drop collection, %s", err)
	}

	// gather collection data for tweets
	data, err := getCollectionData()
	if err != nil {
		return fmt.Errorf("failed to retrieve data for tweet collection, %s", err)
	}

	// creates collection and populates all data
	err = client.InitCollection(data)
	if err != nil {
		return fmt.Errorf("failed to save data in collection, %s", err)
	}

	log.Println("[Info] Database initialization finished")
	return nil
}

// getCollectionData get all tweets data from url
func getCollectionData() ([]interface{}, error) {
	log.Println("[Info] Collecting database tweets data")
	var result []interface{}
	resp, err := http.Get(TweetsCollectionJsonUrl)
	if err != nil {
		return result, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		return result, err
	}
	log.Println("[Info] Collecting database tweets data finished")
	return result, nil
}
