package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/ob-vss-ss19/blatt-4-team1234/userservice/handler"
	"github.com/ob-vss-ss19/blatt-4-team1234/userservice/subscriber"

	example "github.com/ob-vss-ss19/blatt-4-team1234/userservice/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.userservice"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.userservice", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.userservice", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
