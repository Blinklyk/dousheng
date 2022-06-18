package main

import (
	"context"
	"errors"
	"github.com/RaymondCode/simple-demo/initialize"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/pb/rpcUser"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils/rpcdto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type RpcUserService struct {
	rpcUser.UnimplementedRPCUserServiceServer
}

func (*RpcUserService) Register(ctx context.Context, req *rpcUser.RegisterRequest) (*rpcUser.RegisterResponse, error) {
	username := req.Username
	password := req.Password

	// call service
	newUser := &model.User{Name: req.Username, Username: username, Password: password, FollowCount: 0, FollowerCount: 0}
	var userService = service.UserService{}
	err, userReturn := userService.Register(newUser)
	if err != nil {
		return nil, err
	}

	// return
	resp := &rpcUser.RegisterResponse{UserId: userReturn.ID, Token: newUser.Name}
	return resp, nil
}

func (*RpcUserService) Login(ctx context.Context, req *rpcUser.LoginRequest) (*rpcUser.LoginResponse, error) {

	// call service
	user := &model.User{Name: req.Username, Username: req.Username, Password: req.Password}
	var loginService = service.UserService{}
	userReturn, tokenStr, err := loginService.Login(user)
	if tokenStr == "" {
		return nil, errors.New("error tokenStr is empty")
	}
	if err != nil {
		return nil, errors.New("failed: login in")
	}
	resp := &rpcUser.LoginResponse{UserId: userReturn.ID, Token: tokenStr}
	return resp, nil
}

func (*RpcUserService) GetUserInfo(ctx context.Context, req *rpcUser.UserInfoRequest) (*rpcUser.UserInfoResponse, error) {
	// call service
	var checkUserInfoService = service.UserService{}
	returnUser, err := checkUserInfoService.GetUserInfo(req.UserId, req.UserId)
	if err != nil {
		return nil, errors.New("error: db select")
	}

	// to rpc dto
	u := rpcdto.ToUserRpcDTO(returnUser)
	resp := &rpcUser.UserInfoResponse{User: u}
	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	rpcUser.RegisterRPCUserServiceServer(s, &RpcUserService{})

	// init db
	initialize.Init()

	// Register reflection service on gRPC server.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
