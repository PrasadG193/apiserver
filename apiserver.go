package main

import (
	"log"
	"net/http"
)


func init() {
	InitLogger()
	InitRedis()
}

func main() {
	Log.Info("In main")
	log.Fatal(http.ListenAndServe(":12345", SetHandlers()))
}
