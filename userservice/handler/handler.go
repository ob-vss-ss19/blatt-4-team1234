package handler

import (
	"context"

	"github.com/ob-vss-ss19/blatt-4-team1234/commons"

	"github.com/ob-vss-ss19/blatt-4-team1234/reservationservice/proto/reservation"
	"github.com/ob-vss-ss19/blatt-4-team1234/userservice/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	Users map[int64]user.User
}

func (handle *UserHandler) GetAllUsers(ctx context.Context, req *user.GetAllUsersRequest,
	rsp *user.GetAllUsersResponse) error {
	protoUsers := make([]*user.User, len(handle.Users))
	i := 0
	for _, u := range handle.Users {
		u := u
		protoUsers[i] = &u
		i++
	}
	rsp.Users = protoUsers
	return nil
}

func (handle *UserHandler) GetUser(ctx context.Context, req *user.GetUserRequest,
	rsp *user.GetUserResponse) error {
	if err := commons.CheckId(req.Id, "User"); err != nil {
		return err
	}
	u, found := handle.Users[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The User with the ID:%d does not Exist", req.Id)
	}
	rsp.User = &u
	return nil
}

func (handle *UserHandler) RemoveUser(ctx context.Context, req *user.RemoveUserRequest,
	rsp *user.RemoveUserResponse) error {
	if err := commons.CheckId(req.Id, "User"); err != nil {
		return err
	}
	_, found := handle.Users[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The User with the ID:%d does not Exist", req.Id)
	}
	request := reservation.GetReservationsForUserRequest{UserId: req.Id}
	reservationService := reservation.NewReservationService("go.micro.srv.reservationservice", nil)
	response, err := reservationService.GetReservationsForUser(ctx, &request)
	if err != nil {
		return status.Errorf(codes.Internal, "An internal error occurred while getting the reservations for this"+
			" user. Error: "+err.Error())
	}
	if len(response.Reservations) == 0 {
		delete(handle.Users, req.Id)
	} else {
		return status.Errorf(codes.FailedPrecondition, "The User still has open reservations")
	}
	delete(handle.Users, req.Id)
	return nil
}

func (handle *UserHandler) AddUser(ctx context.Context, req *user.AddUserRequest, rsp *user.AddUserResponse) error {
	if req.User == nil {
		return status.Errorf(codes.InvalidArgument, "No User was Provided!")
	}
	if len(req.User.LastName) < 1 {
		return status.Errorf(codes.InvalidArgument, "No LastName was Provided!")
	}
	if len(req.User.FirstName) < 1 {
		return status.Errorf(codes.InvalidArgument, "No FirstName was Provided!")
	}
	if req.User.Age < 1 {
		return status.Errorf(codes.InvalidArgument, "No Age was Provided!")
	}
	req.User.Id = int64(len(handle.Users) + 2)
	handle.Users[int64(len(handle.Users)+2)] = *req.User
	return nil
}

func (handle *UserHandler) InitDB() {
	handle.Users = make(map[int64]user.User)
	handle.Users[1] = user.User{Id: 1, FirstName: "Bob", LastName: "Baumeister", Age: 6}
	handle.Users[2] = user.User{Id: 2, FirstName: "John", LastName: "Wick", Age: 42}
	handle.Users[3] = user.User{Id: 3, FirstName: "Mani", LastName: "Mammut", Age: 17}
	handle.Users[4] = user.User{Id: 4, FirstName: "Jack", LastName: "Sparrow", Age: 31}
}
