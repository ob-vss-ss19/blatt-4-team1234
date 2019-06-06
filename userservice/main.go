package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/ob-vss-ss19/blatt-4-team1234/commons"
	"github.com/ob-vss-ss19/blatt-4-team1234/userservice/handler"

	example "github.com/ob-vss-ss19/blatt-4-team1234/userservice/proto/user"
)

func main() {

	userHandler := new(handler.UserHandler)
	userHandler.InitDB()
	// New Service
	service := micro.NewService(
		micro.Name(commons.GetUserServiceName()),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	err := example.RegisterUserServiceHandler(service.Server(), userHandler)
	if err != nil {
		log.Fatalf("An Error occurred while registering the UserHandler for the Service: %s",
			commons.GetUserServiceName())
	}
	// Run service
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
