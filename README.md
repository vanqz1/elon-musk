# Elon Musk Tweets

## Prerequisites for running

Installed Docker

## Running app

```sh
docker-compose build
```
```sh
docker-compose up -d
```

This will start four containers: 
|  Container name | Address  |  Description | 
|---|---|---|
| tweets_ui  |  localhost:300 |  React application that calles tweets_api to retrieve aggregated tweets data and display it in charts |  
| tweets_api  | localhost:8030  | GoLang app that server aggregated tweets data extracted from mongo | 
|  mongo | localhost:27017  | Container running  MongoDB instance storing tweets data| 
|  mongo_express | localhost:8081  | MongoDB interface - easy for local testing  |



## Service description

The tweets_api service contains one endpoint which aggregates and groups tweets by their creation date based on the  query parameter groupBy.
The allowed values for the query parameter groupBy are: "year", "month", "hour" . 
It is possible to use more than one groupBy query parameter. 

Example 1:
When no query parameter is passed the data returned will be aggregated by year
```sh
http://localhost:8030/api/tweets/aggregate
```
same as
```sh
http://localhost:8030/api/tweets/aggregate?groupBy=year
```


Example 2:
Get aggregated result grouped by each month during the years
```sh
http://localhost:8030/api/tweets/aggregate?groupBy=month&groupBy=year
```

Example 3:
Get aggregated result grouped by time of the day - hour when the tweet was created
```sh
http://localhost:8030/api/tweets/aggregate?groupBy=hour
```




