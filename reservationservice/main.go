package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/ob-vss-ss19/blatt-4-team1234/reservationservice/handler"

	example "github.com/ob-vss-ss19/blatt-4-team1234/reservationservice/proto/reservation"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.reservationservice"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	err := example.RegisterExampleHandler(service.Server(), new(handler.Example))
	if err != nil{
		log.Fatal("An Error occurred while registering the ReservationHandler for the Service: go.micro.src.reservationservice")
	}

	// Run service
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
