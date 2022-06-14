package controller

import (
	"context"
	"encoding/json"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/pb/rpcUser"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/RaymondCode/simple-demo/utils/verify"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {

	// bind rpc request var
	var r request.RegisterRequest
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// rpc init
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := rpcUser.NewRPCUserServiceClient(conn)

	// call server
	resp, err := client.Register(context.Background(), &rpcUser.RegisterRequest{Username: r.Username, Password: r.Password})
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.RegisterResponse{Response: response.Response{StatusCode: -1, StatusMsg: "failed: rpc error " + err.Error()}})
		return
	}

	// return
	token, _ := utils.GenToken(resp.UserId)
	c.JSON(http.StatusOK, response.RegisterResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "success: create register rpcUser",
		},
		UserId: resp.UserId,
		// TODO register token
		Token: token,
	})
}

func Login(c *gin.Context) {

	// bind request var
	var l request.LoginRequest
	if err := c.ShouldBind(&l); err != nil {
		global.App.DY_LOG.Error("bind error", zap.Error(err))
		c.JSON(http.StatusBadRequest, Response{1, "bind error"})
		return
	}

	//verify
	if err := verify.Login(l); err != nil {
		c.JSON(http.StatusBadRequest, Response{1, err.Error()})
		return
	}

	// rpc client
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := rpcUser.NewRPCUserServiceClient(conn)

	// call server
	resp, err := client.Login(context.Background(), &rpcUser.LoginRequest{Username: l.Username, Password: l.Password})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"rpc server error": err.Error()})
		return
	}

	// return
	c.JSON(http.StatusOK, response.LoginResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "success: login in",
		},
		UserId: resp.UserId,
		Token:  resp.Token,
	})

}

// UserInfo get login userInfo from db
func UserInfo(c *gin.Context) {

	// authentication jwt version
	UserStr, _ := c.Get("UserStr")

	var userInfoVar model.User
	if err := json.Unmarshal([]byte(UserStr.(string)), &userInfoVar); err != nil {
		global.App.DY_LOG.Error("session unmarshal error", zap.Error(err))
		c.JSON(http.StatusBadRequest, response.UserInfoResponse{Response: response.Response{StatusCode: 1, StatusMsg: "error: session unmarshal error"}})
		return
	}

	userId := c.Query("user_id")
	userIdNum, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.UserInfoResponse{Response: response.Response{StatusCode: 1, StatusMsg: "error: userID error"}})
	}

	// rpc client
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := rpcUser.NewRPCUserServiceClient(conn)

	// call server
	resp, err := client.GetUserInfo(context.Background(), &rpcUser.UserInfoRequest{UserId: userIdNum})
	returnUser := resp.User

	// DTO
	userinfo := response.UserInfo{
		ID:            returnUser.Id,
		Name:          returnUser.Name,
		FollowCount:   returnUser.FollowCount,
		FollowerCount: returnUser.FollowerCount,
		IsFollow:      returnUser.IsFollow,
		Username:      returnUser.Username,
	}
	c.JSON(http.StatusOK, response.UserInfoResponse{
		Response: response.Response{StatusCode: 0},
		UserInfo: userinfo,
	})
	return

	////session version
	//// get user from redis
	//userID := c.Query("user_id")
	//session := sessions.Default(c)
	//jsonUser := session.Get(userID)
	//log.Println("jsonUser : ", jsonUser)
	//log.Println(c.GetHeader("Cookie"))
	//log.Println(c.GetHeader("Host"))
	//log.Println(c.GetHeader("Connection"))
	//
	////userInfoVar := &userInfoVar{}
	//userInfoVar := &model.User{}
	//
	//err := json.Unmarshal(jsonUser.([]byte), userInfoVar)
	//if err != nil {
	//	c.JSON(http.StatusOK, CheckUserInfoResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "Unmarshal from session failed"},
	//	})
	//	return
	//}
	//
	//if len(userInfoVar.Name) < 0 {
	//	c.JSON(http.StatusOK, CheckUserInfoResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "userName len is 0"},
	//	})
	//	return
	//}
	//
	//
	//c.JSON(http.StatusOK, CheckUserInfoResponse{
	//	Response: Response{StatusCode: 0},
	//	UserInfo: *userInfoVar,
	//})
	//return

}
