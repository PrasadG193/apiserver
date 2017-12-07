package main

import (
	//"fmt"
	"time"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/go-redis/redis"
	)

type resp struct {
	Count	int64
}

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{Addr: "localhost:6379", DB: 0,})
	_, err := redisClient.Ping().Result()
	//logger.Logger.Println(pong, err)
	for err != nil {
		// Wait for redis to come up
		Log.Error(err, ". Wait and retry")
		time.Sleep(5 * time.Second)
		_, err = redisClient.Ping().Result()
	}
	redisClient.FlushDB()
}

func SetHandlers() *mux.Router {
	Log.Info("Set handlers")
	router := mux.NewRouter()
	router.HandleFunc("/login", NewVisitor).Methods("GET")
	return router
}

func NewVisitor(w http.ResponseWriter, req *http.Request) {
	//fmt.Println("New User")
	count, err := redisClient.Incr("client_count").Result()
	if err != nil {
		panic(err)
	}
	Log.Info("Client count :", count)
	r := resp{Count: count}
	json.NewEncoder(w).Encode(r)
}
