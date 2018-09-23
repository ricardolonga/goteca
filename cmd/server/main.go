package main

import (
	"os"

	"os/signal"
	"syscall"

	"github.com/NeowayLabs/logger"
	"github.com/ricardolonga/goteca"
	"github.com/ricardolonga/goteca/http"
	"github.com/ricardolonga/goteca/mongo"
	"gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		logger.Fatal("error on Mongo connection: %q", err)
	}

	repository := mongo.NewDao(session)

	service := goteca.NewService(repository)

	handler := http.NewHandler(service)

	server := http.NewServer("8080", handler)

	server.ListenAndServe()

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan

	server.Shutdown()
}
