package controller

import (
	"context"
	"encoding/json"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/pb/rpcFollow"
	"github.com/RaymondCode/simple-demo/utils/respToDTO"
	"github.com/RaymondCode/simple-demo/utils/verify"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"strconv"
)

// RelationAction actionType = 1: follow; actionType = 2: cancel follow
func RelationAction(c *gin.Context) {
	// authorization
	UserStr, _ := c.Get("UserStr")
	log.Println("UserStr: ", UserStr)

	var userInfoVar model.User
	if err := json.Unmarshal([]byte(UserStr.(string)), &userInfoVar); err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "error: session unmarshal error"})
		return
	}

	// bind
	var relationActionRequest request.RelationActionRequest
	if err := c.ShouldBind(&relationActionRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "bind error"})
		return
	}

	// verify
	if err := verify.Relation(relationActionRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{1, err.Error()})
		return
	}

	// cannot follow myself
	if relationActionRequest.ToUserID == strconv.Itoa(int(userInfoVar.ID)) {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "cannot follow myself"})
		return
	}

	// rpc client
	conn, err := grpc.Dial("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := rpcFollow.NewRPCFollowServiceClient(conn)

	// call server
	resp, err := client.FollowAction(context.Background(), &rpcFollow.FollowActionReq{
		Token:      relationActionRequest.Token,
		ToUserId:   relationActionRequest.ToUserID,
		ActionType: relationActionRequest.ActionType,
		UserId:     userInfoVar.ID,
	})
	if err != nil {
		global.App.DY_LOG.Error("rpc error", zap.Error(err))
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: resp.StatusCode})

}

// FollowList get follow list of the current user
func FollowList(c *gin.Context) {
	// authorization
	UserStr, _ := c.Get("UserStr")
	log.Println("UserStr: ", UserStr)

	var userInfoVar model.User
	if err := json.Unmarshal([]byte(UserStr.(string)), &userInfoVar); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error: session unmarshal error"})
		return
	}
	// bind var
	var followListRequest request.FollowListRequest
	if err := c.ShouldBind(&followListRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "bind error"})
		return
	}
	// rpc client
	conn, err := grpc.Dial("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := rpcFollow.NewRPCFollowServiceClient(conn)

	// call server
	resp, err := client.GetFollowList(context.Background(), &rpcFollow.FollowListReq{
		UserId: followListRequest.UserID,
		Token:  followListRequest.Token,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	followList := respToDTO.GetUserListDTO(resp.UserList)
	c.JSON(http.StatusOK, response.FollowListResponse{
		Response: response.Response{
			StatusCode: 0,
		},
		UserList: *followList,
	})
}

// FollowerList get follower list of the current user
func FollowerList(c *gin.Context) {

	UserStr, _ := c.Get("UserStr")

	var userInfoVar model.User
	if err := json.Unmarshal([]byte(UserStr.(string)), &userInfoVar); err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error: session unmarshal error"})
		return
	}
	// bind var
	var followerListRequest request.FollowerListRequest
	if err := c.ShouldBind(&followerListRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "bind error"})
		return
	}
	// rpc client
	conn, err := grpc.Dial("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := rpcFollow.NewRPCFollowServiceClient(conn)

	// call server
	resp, err := client.GetFollowerList(context.Background(), &rpcFollow.FollowerListReq{
		UserId: followerListRequest.UserID,
		Token:  followerListRequest.Token,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	followerList := respToDTO.GetUserListDTO(resp.UserList)
	c.JSON(http.StatusOK, response.FollowListResponse{
		Response: response.Response{
			StatusCode: 0,
		},
		UserList: *followerList,
	})

}
