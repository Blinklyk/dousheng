package main

import (
	"context"
	"errors"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/initialize"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/pb/rpcVideo"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils/rpcdto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type RpcVideoService struct {
	rpcVideo.UnimplementedRPCVideoServiceServer
}

func (*RpcVideoService) GetPublishList(ctx context.Context, req *rpcVideo.PublishListRequest) (*rpcVideo.PublishListResponse, error) {

	// call service
	ps := service.PublishService{}
	publishVideos, err := ps.PublishList(&request.PublishListRequest{Token: req.Token, UserID: req.UserId})
	if err != nil {
		return nil, errors.New("error in publish service: " + err.Error())
	}

	// to rpc dto
	videos := rpcdto.ToVideoListRpcDTO(publishVideos)
	resp := &rpcVideo.PublishListResponse{VideoList: videos}
	return resp, nil
}

func (*RpcVideoService) Feed(ctx context.Context, req *rpcVideo.FeedReq) (*rpcVideo.FeedResp, error) {
	l := req.LatestTime
	token := req.Token
	resp := &rpcVideo.FeedResp{}
	// call service
	fs := service.FeedService{}
	// if request doesn't contain token
	if token == "" {
		feedList, err := fs.FeedWithoutToken()
		if err != nil {
			// 判断是否是超时错误
			statusErr, ok := status.FromError(err)
			if ok {
				if statusErr.Code() == codes.DeadlineExceeded {
					global.App.DY_LOG.Error("client.Search err: deadline", zap.Error(err))
					log.Println("client.Search err: deadline ", zap.Error(err))
				}
			}
			global.App.DY_LOG.Error("client.Search err: ", zap.Error(err))
			log.Println("client.Search err: ", zap.Error(err))

			return nil, errors.New("error:feed without token" + err.Error())
		}
		feedListRpcDTO := rpcdto.ToVideoListRpcDTO(*feedList)
		resp.VideoList = feedListRpcDTO
	}

	// if request contains token
	if token != "" {
		feedList, err := fs.FeedWithToken(&request.FeedRequest{
			LatestTime: l,
			Token:      token,
		})
		if err != nil {
			return nil, errors.New("error:feed with token" + err.Error())
		}
		feedListRpcDTO := rpcdto.ToVideoListRpcDTO(*feedList)
		resp.VideoList = feedListRpcDTO
	}
	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	rpcVideo.RegisterRPCVideoServiceServer(s, &RpcVideoService{})

	// init db
	initialize.Init()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
