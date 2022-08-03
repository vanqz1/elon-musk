package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"tweets/storage"
	"tweets/storage/mongo/aggregations"
)

const MongoDbURI = "mongodb://mongo:27017"

var (
	once           sync.Once
	clientInstance *mongo.Client
	instanceError  error
)

type mongoDbClient struct {
	client *mongo.Client
	config DbConfig
}

func NewDbMongoClient(config DbConfig) (storage.IDbClient, error) {
	client, err := getMongoClient()
	if err != nil {
		return nil, err
	}
	return &mongoDbClient{client: client, config: config}, nil
}

func (mc *mongoDbClient) DropCollection() error {
	err := mc.client.Database(mc.config.DatabaseName).Collection(mc.config.CollectionName).Drop(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

func (mc *mongoDbClient) InitCollection(data []interface{}) error {
	collection := mc.client.Database(mc.config.DatabaseName).Collection(mc.config.CollectionName)
	_, err := collection.InsertMany(context.TODO(), data)
	if err != nil {
		return err
	}

	return nil
}

func (mc *mongoDbClient) Aggregate(pipelineBuilder aggregations.IPipelinesBuilder) ([]interface{}, error) {
	var result []interface{}

	aggregationPipeline := pipelineBuilder.Build()
	collection := mc.client.Database(mc.config.DatabaseName).Collection(mc.config.CollectionName)

	cur, err := collection.Aggregate(context.TODO(), aggregationPipeline)
	if err != nil {
		return result, err
	}

	for cur.Next(context.TODO()) {
		var elem bson.M
		err := cur.Decode(&elem)
		if err != nil {
			return result, err
		}

		result = append(result, elem)
	}

	if err := cur.Err(); err != nil {
		return result, err
	}

	cur.Close(context.TODO())

	return result, err
}

func getMongoClient() (*mongo.Client, error) {
	once.Do(func() {
		log.Println("[Info] Connecting to mongodb")
		clientOptions := options.Client().ApplyURI(MongoDbURI)
		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			instanceError = err
		}

		err = client.Ping(context.TODO(), nil)

		if err != nil {
			instanceError = err
		}

		log.Println("[Info] Successfully connected to mongodb")
		clientInstance = client
	})

	return clientInstance, instanceError
}
