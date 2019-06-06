package commons

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// nolint:gochecknoglobals
	userService = "go.micro.srv.userservice"
	// nolint:gochecknoglobals
	hallService = "go.micro.srv.hallservice"
	// nolint:gochecknoglobals
	reservationService = "go.micro.srv.reservationservice"
	// nolint:gochecknoglobals
	movieService = "go.micro.srv.movieservice"
	// nolint:gochecknoglobals
	showService = "go.micro.srv.showservice"
)

func CheckId(id int64, idName string) error {
	if id <= 0 {
		return status.Errorf(codes.InvalidArgument, "No Valid %s-Id was Submitted! ID <= 0", idName)
	}
	return nil
}
