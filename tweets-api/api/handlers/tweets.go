package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"tweets/api/services"
	"tweets/api/services/enums"
)

const GroupByQueryParameter = "groupBy"

type TweetsHandler struct {
	tweetsService *services.TweetsService
}

// NewTweetsHandler creates new instance of TweetsHandler
func NewTweetsHandler(service *services.TweetsService) *TweetsHandler {
	return &TweetsHandler{
		tweetsService: service,
	}
}

// TweetsAggregationHandler is the handler for tweets aggregations service
func (th *TweetsHandler) TweetsAggregationHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		r := recover()
		if r != nil {
			var err error
			switch t := r.(type) {
			case string:
				err = errors.New(t)
			case error:
				err = t
			default:
				err = errors.New("unknown error")
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	queries := r.URL.Query()
	groupByQuery := queries[GroupByQueryParameter]

	groupByValues, err := th.parseGroupByQueryParams(groupByQuery)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s: invalid value for groupBy query parameter. Valid values: %s, %s, %s",
			http.StatusText(400), enums.Year, enums.Hour, enums.Month), 400)
		return
	}

	result, err := th.tweetsService.GetAggregatedResult(groupByValues)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	responseBytes, _ := json.Marshal(result)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Write(responseBytes)
}

// parseGroupByQueryParams parse the query parameters
func (th *TweetsHandler) parseGroupByQueryParams(groupByQuery []string) ([]enums.GroupByEnum, error) {
	var result []enums.GroupByEnum

	if len(groupByQuery) != 0 {
		var err error
		result, err = enums.ConvertToGroupByArray(groupByQuery)
		if err != nil {
			return nil, err
		}
	} else {
		result = append(result, enums.Year)
	}

	return result, nil
}
