package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/ob-vss-ss19/blatt-4-team1234/commons"
	"github.com/ob-vss-ss19/blatt-4-team1234/reservationservice/handler"

	example "github.com/ob-vss-ss19/blatt-4-team1234/reservationservice/proto/reservation"
)

func main() {

	reservationHandler := new(handler.ReservationHandler)
	reservationHandler.InitDB()
	// New Service
	service := micro.NewService(
		micro.Name(commons.GetReservationServiceName()),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	err := example.RegisterReservationServiceHandler(service.Server(), reservationHandler)
	if err != nil {
		log.Fatalf("An Error occurred while registering the ReservationHandler"+
			" for the Service: %s", commons.GetReservationServiceName())
	}

	// Run service
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
