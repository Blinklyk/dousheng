package main

import (
	"context"
	"github.com/RaymondCode/simple-demo/initialize"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/pb/rpcFollow"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils/rpcdto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type RpcFollowService struct {
	rpcFollow.UnimplementedRPCFollowServiceServer
}

func (*RpcFollowService) FollowAction(ctx context.Context, req *rpcFollow.FollowActionReq) (*rpcFollow.FollowActionResp, error) {
	userId := req.UserId

	// call the service
	relationService := &service.RelationService{}
	if err := relationService.RelationAction(userId, &request.RelationActionRequest{
		Token:      req.Token,
		ToUserID:   req.ToUserId,
		ActionType: req.ActionType,
	}); err != nil {
		return nil, err
	}

	// return rpc resp
	resp := &rpcFollow.FollowActionResp{StatusCode: 0}
	return resp, nil
}

func (*RpcFollowService) GetFollowList(ctx context.Context, req *rpcFollow.FollowListReq) (*rpcFollow.FollowListResp, error) {
	// call service
	relationService := &service.RelationService{}
	followList, err := relationService.FollowList(&request.FollowListRequest{
		Token:  req.Token,
		UserID: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	followListRpcDTO := rpcdto.ToUserListDTO(&followList)
	resp := rpcFollow.FollowListResp{StatusCode: 0, UserList: followListRpcDTO}
	return &resp, nil
}

func (*RpcFollowService) GetFollowerList(ctx context.Context, req *rpcFollow.FollowerListReq) (*rpcFollow.FollowerListResp, error) {
	// call service
	relationService := &service.RelationService{}
	followerList, err := relationService.FollowerList(&request.FollowerListRequest{
		Token:  req.Token,
		UserID: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	followerListRpcDTO := rpcdto.ToUserListDTO(&followerList)
	resp := rpcFollow.FollowerListResp{StatusCode: 0, UserList: followerListRpcDTO}
	return &resp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	rpcFollow.RegisterRPCFollowServiceServer(s, &RpcFollowService{})

	// init db
	initialize.Init()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
