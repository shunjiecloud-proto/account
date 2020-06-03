// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: account.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Account service

func NewAccountEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Account service

type AccountService interface {
	//  创建用户
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...client.CallOption) (*CreateUserResponse, error)
	//  获取用户信息
	UserInfo(ctx context.Context, in *UserInfoRequest, opts ...client.CallOption) (*UserInfoResponse, error)
	//  检查用户是否可以被注册
	RegistrationCheck(ctx context.Context, in *RegistrationCheckRequest, opts ...client.CallOption) (*RegistrationCheckResponse, error)
	//  登陆
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	//  登出
	Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error)
}

type accountService struct {
	c    client.Client
	name string
}

func NewAccountService(name string, c client.Client) AccountService {
	return &accountService{
		c:    c,
		name: name,
	}
}

func (c *accountService) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...client.CallOption) (*CreateUserResponse, error) {
	req := c.c.NewRequest(c.name, "Account.CreateUser", in)
	out := new(CreateUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...client.CallOption) (*UserInfoResponse, error) {
	req := c.c.NewRequest(c.name, "Account.UserInfo", in)
	out := new(UserInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) RegistrationCheck(ctx context.Context, in *RegistrationCheckRequest, opts ...client.CallOption) (*RegistrationCheckResponse, error) {
	req := c.c.NewRequest(c.name, "Account.RegistrationCheck", in)
	out := new(RegistrationCheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Account.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error) {
	req := c.c.NewRequest(c.name, "Account.Logout", in)
	out := new(LogoutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountHandler interface {
	//  创建用户
	CreateUser(context.Context, *CreateUserRequest, *CreateUserResponse) error
	//  获取用户信息
	UserInfo(context.Context, *UserInfoRequest, *UserInfoResponse) error
	//  检查用户是否可以被注册
	RegistrationCheck(context.Context, *RegistrationCheckRequest, *RegistrationCheckResponse) error
	//  登陆
	Login(context.Context, *LoginRequest, *LoginResponse) error
	//  登出
	Logout(context.Context, *LogoutRequest, *LogoutResponse) error
}

func RegisterAccountHandler(s server.Server, hdlr AccountHandler, opts ...server.HandlerOption) error {
	type account interface {
		CreateUser(ctx context.Context, in *CreateUserRequest, out *CreateUserResponse) error
		UserInfo(ctx context.Context, in *UserInfoRequest, out *UserInfoResponse) error
		RegistrationCheck(ctx context.Context, in *RegistrationCheckRequest, out *RegistrationCheckResponse) error
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		Logout(ctx context.Context, in *LogoutRequest, out *LogoutResponse) error
	}
	type Account struct {
		account
	}
	h := &accountHandler{hdlr}
	return s.Handle(s.NewHandler(&Account{h}, opts...))
}

type accountHandler struct {
	AccountHandler
}

func (h *accountHandler) CreateUser(ctx context.Context, in *CreateUserRequest, out *CreateUserResponse) error {
	return h.AccountHandler.CreateUser(ctx, in, out)
}

func (h *accountHandler) UserInfo(ctx context.Context, in *UserInfoRequest, out *UserInfoResponse) error {
	return h.AccountHandler.UserInfo(ctx, in, out)
}

func (h *accountHandler) RegistrationCheck(ctx context.Context, in *RegistrationCheckRequest, out *RegistrationCheckResponse) error {
	return h.AccountHandler.RegistrationCheck(ctx, in, out)
}

func (h *accountHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.AccountHandler.Login(ctx, in, out)
}

func (h *accountHandler) Logout(ctx context.Context, in *LogoutRequest, out *LogoutResponse) error {
	return h.AccountHandler.Logout(ctx, in, out)
}
