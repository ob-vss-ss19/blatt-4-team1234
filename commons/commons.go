package commons

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetUserServiceName() string {
	return "go.micro.srv.userservice"
}
func GetHallServiceName() string {
	return "go.micro.srv.hallservice"
}
func GetShowServiceName() string {
	return "go.micro.srv.showservice"
}
func GetMovieServiceName() string {
	return "go.micro.srv.movieservice"
}
func GetReservationServiceName() string {
	return "go.micro.srv.reservationservice"
}

func CheckId(id int64, idName string) error {
	if id <= 0 {
		return status.Errorf(codes.InvalidArgument, "No Valid %s-Id was Submitted! ID <= 0", idName)
	}
	return nil
}
