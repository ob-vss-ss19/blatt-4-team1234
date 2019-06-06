package commons

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CheckId(id int64, idName string) error {
	if id <= 0 {
		return status.Errorf(codes.InvalidArgument, "No Valid %s-Id was Submitted! ID <= 0", idName)
	}
	return nil
}
