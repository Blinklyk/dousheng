package main

import (
	"context"
	"errors"
	"github.com/RaymondCode/simple-demo/initialize"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/pb/rpcComment"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils/rpcdto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type RpcCommentService struct {
	rpcComment.UnimplementedRPCCommentServiceServer

}

func (*RpcCommentService) Comment(ctx context.Context, req *rpcComment.CommentRequest) (*rpcComment.CommentResponse, error) {

	actionType := req.ActionType
	userId := req.UserId
	c := &request.CommentRequest{
		Token:       req.Token,
		VideoID:     req.VideoId,
		ActionType:  req.ActionType,
		CommentText: req.CommentText,
		CommentID:   req.CommentId,
	}

	cs := service.CommentService{}
	// add comment action
	if actionType == "1" {
		commentVar, err := cs.CommentAction(userId, c)
		if err != nil {
			return nil, err
		}
		CommentRpcDTO := rpcdto.ToCommentRpcDTO(commentVar)
		resp := &rpcComment.CommentResponse{StatusCode: 0, Comment: CommentRpcDTO}
		return resp, nil
	}

	// delete comment action
	if actionType == "2" {
		if err := cs.DeleteCommentAction(c); err != nil {
			return nil, err
		}
		resp := &rpcComment.CommentResponse{StatusCode: 0}
		return resp, nil
	}
	return nil, errors.New("error action type")
}

func (*RpcCommentService) GetCommentList(ctx context.Context, req *rpcComment.CommentListReq) (*rpcComment.CommentListResp, error) {
	userID := req.UserId
	// call service
	cs := service.CommentService{}
	commentList, err := cs.CommentList(userID, &request.CommentListRequest{
		Token:   req.Token,
		VideoID: req.VideoId,
	})
	if err != nil {
		return nil, err
	}
	c := rpcdto.ToCommentListRpcDTO(*commentList)
	resp := rpcComment.CommentListResp{StatusCode: 0, CommentList: c}
	return &resp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	rpcComment.RegisterRPCCommentServiceServer(s, &RpcCommentService{})

	// init db
	initialize.Init()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
