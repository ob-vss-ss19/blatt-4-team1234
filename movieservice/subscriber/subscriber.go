package subscriber

import (
	"context"
	"github.com/micro/go-log"

	example "github.com/ob-vss-ss19/blatt-4-team1234/movieservice/proto/movie"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *example.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
