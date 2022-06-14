package main

import (
	"context"
	"errors"
	"github.com/RaymondCode/simple-demo/initialize"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/pb/rpcFavorite"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils/rpcdto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type RpcFavoriteService struct {
	rpcFavorite.UnimplementedRPCFavoriteServiceServer
}

func (*RpcFavoriteService) FavoriteAction(ctx context.Context, req *rpcFavorite.FavoriteRequest) (*rpcFavorite.FavoriteResponse, error) {
	userId := req.UserId
	favoriteRequest := &request.FavoriteRequest{Token: req.Token, VideoID: req.VideoId, ActionType: req.ActionType}
	// call service
	fs := service.FavoriteService{}
	err := fs.FavoriteAction(userId, favoriteRequest)
	if err != nil {
		return nil, errors.New("error in favorite action service: " + err.Error())
	}

	// resp
	resp := &rpcFavorite.FavoriteResponse{StatusCode: 0}
	return resp, nil
}

func (*RpcFavoriteService) FavoriteList(ctx context.Context, req *rpcFavorite.FavoriteListRequest) (*rpcFavorite.FavoriteListResponse, error) {
	favoriteListRequest := &request.FavoriteListRequest{UserID: req.UserId, Token: req.Token}
	// call service
	fs := service.FavoriteService{}
	favoriteVideoList, err := fs.FavoriteList(favoriteListRequest)
	if err != nil {
		return &rpcFavorite.FavoriteListResponse{StatusCode: 1, StatusMsg: error.Error(err)}, nil
	}
	videos := rpcdto.ToVideoListRpcDTO(*favoriteVideoList)
	resp := &rpcFavorite.FavoriteListResponse{StatusCode: 0, VideoList: videos}
	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	rpcFavorite.RegisterRPCFavoriteServiceServer(s, &RpcFavoriteService{})

	// init db
	initialize.Init()

	// Register reflection service on gRPC server.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
