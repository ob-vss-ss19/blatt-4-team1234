package main

import (
	"log"

	"github.com/ob-vss-ss19/blatt-4-team1234/commons"

	"github.com/micro/go-micro"
	"github.com/ob-vss-ss19/blatt-4-team1234/hallservice/handler"
	proto "github.com/ob-vss-ss19/blatt-4-team1234/hallservice/proto/hall"
)

func main() {

	hallHandler := new(handler.HallHandler)
	hallHandler.InitDB()

	// New Service
	service := micro.NewService(
		micro.Name(commons.GetHallServiceName()),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	err := proto.RegisterHallServiceHandler(service.Server(), hallHandler)
	if err != nil {
		log.Fatalf("An Error occurred while registering the HallHandler for the Service: %s",
			commons.GetHallServiceName())
	}

	// Run service
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
