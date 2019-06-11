package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/ob-vss-ss19/blatt-4-team1234/commons"
	"github.com/ob-vss-ss19/blatt-4-team1234/showservice/handler"

	example "github.com/ob-vss-ss19/blatt-4-team1234/showservice/proto/show"
)

func main() {

	showHandler := handler.InitDB()
	// New Service
	service := micro.NewService(
		micro.Name(commons.GetShowServiceName()),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	err := example.RegisterShowServiceHandler(service.Server(), showHandler)
	if err != nil {
		log.Fatalf("An Error occurred while registering the ShowHandler for the Service: %s",
			commons.GetShowServiceName())
	}
	// Run service
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
