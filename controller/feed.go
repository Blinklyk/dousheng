package controller

import (
	"context"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/pb/rpcVideo"
	"github.com/RaymondCode/simple-demo/utils/respToDTO"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"time"
)

// Feed token is optional here
func Feed(c *gin.Context) {

	// bind request var
	var feedRequest request.FeedRequest
	if err := c.ShouldBind(&feedRequest); err != nil {
		global.App.DY_LOG.Error("feed bind error!", zap.Error(err))
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "feed bind error"})
		return
	}

	// rpc client
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		global.App.DY_LOG.Error("client conn error!", zap.Error(err))
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := rpcVideo.NewRPCVideoServiceClient(conn)
	// set conn deadline
	route(context.Background(), 2)

	// call server
	resp, err := client.Feed(context.Background(), &rpcVideo.FeedReq{
		LatestTime: feedRequest.LatestTime,
		Token:      feedRequest.Token,
	})
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
		global.App.DY_LOG.Error("feed server error!", zap.Error(err))
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	feedListReturn := respToDTO.GetVideoListDTo(resp.VideoList)
	if feedListReturn == nil {
		global.App.DY_LOG.Info("get feedList null")
	}
	c.JSON(http.StatusOK, response.FeedResponse{
		Response:  response.Response{StatusCode: 0},
		VideoList: feedListReturn,
		NextTime:  time.Now().Unix(),
	})
	return
}

func route(ctx context.Context, i time.Duration) {
	clientDeadline := time.Now().Add(time.Duration(i * time.Second))
	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()

}
