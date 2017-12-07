package main

import (
	"log"
	"net/http"
	"github.com/PrasadG193/apiserver/handler"
	"github.com/PrasadG193/apiserver/logger"
	)


func init() {
	logger.InitLogger()
	handler.InitRedis()
}

func main() {
	logger.Log.Info("In main")
	log.Fatal(http.ListenAndServe(":12345", handler.SetHandlers()))
}
