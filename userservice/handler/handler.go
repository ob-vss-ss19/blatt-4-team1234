package handler

import (
	"context"
	"github.com/ob-vss-ss19/blatt-4-team1234/userservice/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct{
	Users map[int64]User
}



type User struct {
	FirstName string
	LastName string
	Age int64
}

func (handle *UserHandler) GetAllUsers(ctx context.Context, req *user.GetAllUsersRequest, rsp *user.GetAllUsersResponse) error {
	var protoUsers []*user.User
	for i, u := range handle.Users {
		protoUsers = append(protoUsers, &user.User{Id:i,FirstName: u.FirstName,LastName: u.LastName,Age: u.Age})
	}
	rsp.Users = protoUsers
	return nil
}

func (handle *UserHandler) GetUser(ctx context.Context,req *user.GetUserRequest,rsp *user.GetUserResponse) error {
	u,found := handle.Users[req.Id]
	if !found{
		return status.Errorf(codes.NotFound,"The User with the ID:%d does not Exist",req.Id)
	}
	rsp.User = &user.User{Id:req.Id,FirstName:u.FirstName,LastName:u.LastName,Age:u.Age}
	return nil
}

func (handle *UserHandler) RemoveUser(ctx context.Context, req *user.RemoveUserRequest, rsp *user.RemoveUserResponse) error{
	_,found := handle.Users[req.Id]
	if !found{
		return status.Errorf(codes.NotFound,"The User with the ID:%d does not Exist",req.Id)
	}
	delete(handle.Users,req.Id)
	return nil
}

func (handle *UserHandler) AddUser(ctx context.Context,req *user.AddUserRequest,rsp *user.AddUserResponse) error {
	handle.Users[int64(len(handle.Users)+1)] = User{FirstName:req.User.FirstName,LastName:req.User.LastName,Age:req.User.Age}
	return nil
}

func (handle *UserHandler) InitDB(){
	handle.Users = make(map[int64]User)
	handle.Users[0]= User{FirstName:"Bob",LastName:"Baumeister",Age:6}
	handle.Users[1]= User{FirstName:"John",LastName:"Wick",Age:42}
	handle.Users[2]= User{FirstName:"Mani",LastName:"Mammut",Age:17}
	handle.Users[3]= User{FirstName:"Jack",LastName:"Sparrow",Age:31}
}
