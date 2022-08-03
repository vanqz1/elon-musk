package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tweets/api/handlers"
	"tweets/api/services"
)

const port = ":8030"

type IService interface {
	Start() error
}

type service struct {
}

// NewService creates new instance of service
func NewService() IService {
	return &service{}
}

// Start starts the service
func (s *service) Start() error {
	log.Println("[Info] Service starting")
	err := services.InitDatabase()
	if err != nil {
		return err
	}

	tweetsService, err := services.NewTweetsService()
	if err != nil {
		return err
	}

	log.Println("[Info] Listening...")
	tweetsHandler := handlers.NewTweetsHandler(tweetsService)
	r := mux.NewRouter()

	r.HandleFunc("/api/tweets/aggregate", tweetsHandler.TweetsAggregationHandler)

	err = http.ListenAndServe(port, r)
	if err != nil {
		return err
	}

	return nil
}
