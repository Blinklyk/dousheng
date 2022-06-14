package controller

import (
	"context"
	"encoding/json"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/pb/rpcComment"
	"github.com/RaymondCode/simple-demo/utils/verify"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

// CommentAction add comment or delete comment
func CommentAction(c *gin.Context) {

	// authentication
	UserStr, _ := c.Get("UserStr")

	var userInfoVar model.User
	if err := json.Unmarshal([]byte(UserStr.(string)), &userInfoVar); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 1, StatusMsg: "error: session unmarshal error"})
		return
	}

	// bind request var
	var commentRequest request.CommentRequest
	if err := c.ShouldBind(&commentRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "bind error " + err.Error()})
		return
	}

	// verify
	if err := verify.Comment(commentRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 1, StatusMsg: "非法数据 "})
		return
	}

	// rpc client
	conn, err := grpc.Dial("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := rpcComment.NewRPCCommentServiceClient(conn)

	// call server
	resp, err := client.Comment(context.Background(), &rpcComment.CommentRequest{
		Token:       commentRequest.Token,
		VideoId:     commentRequest.VideoID,
		ActionType:  commentRequest.ActionType,
		CommentText: commentRequest.CommentText,
		CommentId:   commentRequest.CommentID,
		UserId:      userInfoVar.ID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error in commentAction: " + err.Error()})
	}
	c.JSON(http.StatusOK, response.CommentActionResponse{
		Response: response.Response{StatusCode: 0},
		Comment:  resp.Comment,
	})

	// call service: action_type = 1 add comment; action_type = 2 delete comment

}

// CommentList get comments list of a video
func CommentList(c *gin.Context) {

	// authentication
	UserStr, _ := c.Get("UserStr")
	var userInfoVar model.User
	if err := json.Unmarshal([]byte(UserStr.(string)), &userInfoVar); err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "error: session unmarshal error"})
		return
	}

	// bind request var
	var commentListRequest request.CommentListRequest
	if err := c.ShouldBind(&commentListRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "bind error "})
		return
	}

	// rpc client
	conn, err := grpc.Dial("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := rpcComment.NewRPCCommentServiceClient(conn)

	// call server
	resp, err := client.GetCommentList(context.Background(), &rpcComment.CommentListReq{
		Token:   commentListRequest.Token,
		VideoId: commentListRequest.VideoID,
		UserId:  userInfoVar.ID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error in commentList: " + err.Error()})
		return
	}

	// return
	c.JSON(http.StatusOK, response.CommentListResponse{
		Response:    response.Response{StatusCode: 0},
		CommentList: resp.CommentList,
	})
}
