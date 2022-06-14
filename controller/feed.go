package controller

import (
	"context"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/pb/rpcVideo"
	"github.com/RaymondCode/simple-demo/utils/respToDTO"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"
)

// Feed token is optional here
func Feed(c *gin.Context) {

	// bind request var
	var feedRequest request.FeedRequest
	if err := c.ShouldBind(&feedRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "feed bind error"})
	}

	// rpc client
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := rpcVideo.NewRPCVideoServiceClient(conn)

	// call server
	resq, err := client.Feed(context.Background(), &rpcVideo.FeedReq{
		LatestTime: feedRequest.LatestTime,
		Token:      feedRequest.Token,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	feedListReturn := respToDTO.GetVideoListDTo(resq.VideoList)
	c.JSON(http.StatusOK, response.FeedResponse{
		Response:  response.Response{StatusCode: 0},
		VideoList: feedListReturn,
		NextTime:  time.Now().Unix(),
	})
	return
}
