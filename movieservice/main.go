package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/ob-vss-ss19/blatt-4-team1234/movieservice/handler"
	example "github.com/ob-vss-ss19/blatt-4-team1234/movieservice/proto/movie"
)

func main() {

	movieHandler := new(handler.MovieHandler)
	movieHandler.InitDB()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.movieservice"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	err := example.RegisterMovieServiceHandler(service.Server(), movieHandler)
	if err != nil{
		log.Fatal("An Error occurred while registering the MovieHandler for the Service: go.micro.src.movieservice")
	}

	// Run service
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
