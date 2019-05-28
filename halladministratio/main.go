package halladministration

import (
	"context"
	"fmt"

	"github.com/ob-vss-ss19/blatt-4-team1234/messages"

	micro "github.com/micro/go-micro"
)

type HallAdministrator struct{}

func (g *HallAdministrator) Hello(ctx context.Context, req *messages.HallRequest, rsp *messages.HallResponse) error {
	haals := GetAllHaals()
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("haaladministrator"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
