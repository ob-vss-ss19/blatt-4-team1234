// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: userservice/proto/user/user.proto

package user

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

import (
	context "context"

	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserService interface {
	GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...client.CallOption) (*GetAllUsersResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error)
	AddUser(ctx context.Context, in *AddUserRequest, opts ...client.CallOption) (*AddUserResponse, error)
	RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...client.CallOption) (*RemoveUserResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...client.CallOption) (*GetAllUsersResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetAllUsers", in)
	out := new(GetAllUsersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetUser", in)
	out := new(GetUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AddUser(ctx context.Context, in *AddUserRequest, opts ...client.CallOption) (*AddUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.AddUser", in)
	out := new(AddUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...client.CallOption) (*RemoveUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.RemoveUser", in)
	out := new(RemoveUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	GetAllUsers(context.Context, *GetAllUsersRequest, *GetAllUsersResponse) error
	GetUser(context.Context, *GetUserRequest, *GetUserResponse) error
	AddUser(context.Context, *AddUserRequest, *AddUserResponse) error
	RemoveUser(context.Context, *RemoveUserRequest, *RemoveUserResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		GetAllUsers(ctx context.Context, in *GetAllUsersRequest, out *GetAllUsersResponse) error
		GetUser(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error
		AddUser(ctx context.Context, in *AddUserRequest, out *AddUserResponse) error
		RemoveUser(ctx context.Context, in *RemoveUserRequest, out *RemoveUserResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) GetAllUsers(ctx context.Context, in *GetAllUsersRequest, out *GetAllUsersResponse) error {
	return h.UserServiceHandler.GetAllUsers(ctx, in, out)
}

func (h *userServiceHandler) GetUser(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error {
	return h.UserServiceHandler.GetUser(ctx, in, out)
}

func (h *userServiceHandler) AddUser(ctx context.Context, in *AddUserRequest, out *AddUserResponse) error {
	return h.UserServiceHandler.AddUser(ctx, in, out)
}

func (h *userServiceHandler) RemoveUser(ctx context.Context, in *RemoveUserRequest, out *RemoveUserResponse) error {
	return h.UserServiceHandler.RemoveUser(ctx, in, out)
}
