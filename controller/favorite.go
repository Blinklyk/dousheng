package controller

import (
	"context"
	"encoding/json"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/pb/rpcFavorite"
	"github.com/RaymondCode/simple-demo/utils/respToDTO"
	"github.com/RaymondCode/simple-demo/utils/verify"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

// FavoriteAction directly update db
func FavoriteAction(c *gin.Context) {

	// authentication
	UserStr, _ := c.Get("UserStr")

	var userInfoVar model.User
	if err := json.Unmarshal([]byte(UserStr.(string)), &userInfoVar); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "error: session unmarshal error"})
		return
	}

	// bind request var
	var favoriteRequest request.FavoriteRequest
	if err := c.ShouldBind(&favoriteRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "bind error "})
		return
	}

	//verify
	if err := verify.Favorite(favoriteRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{1, "非法数据"})
		return
	}

	// rpc client
	conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := rpcFavorite.NewRPCFavoriteServiceClient(conn)

	// call server
	resp, err := client.FavoriteAction(context.Background(), &rpcFavorite.FavoriteRequest{UserId: userInfoVar.ID, Token: favoriteRequest.Token, VideoId: favoriteRequest.VideoID, ActionType: favoriteRequest.ActionType})
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "error in favorite action service: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{StatusCode: resp.StatusCode})
}

// FavoriteList get from favorite table
func FavoriteList(c *gin.Context) {

	// authentication
	UserStr, _ := c.Get("UserStr")

	var userInfoVar model.User
	if err := json.Unmarshal([]byte(UserStr.(string)), &userInfoVar); err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "error: session unmarshal error"})
		return
	}

	// bind request var
	var favoriteListRequest request.FavoriteListRequest
	if err := c.ShouldBind(&favoriteListRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "bind error "})
		return
	}

	// rpc client
	conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := rpcFavorite.NewRPCFavoriteServiceClient(conn)

	// call server
	resp, err := client.FavoriteList(context.Background(), &rpcFavorite.FavoriteListRequest{UserId: favoriteListRequest.UserID, Token: favoriteListRequest.Token})
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error: favoriteList service" + err.Error()})
		return
	}
	favoriteList := respToDTO.GetVideoListDTo(resp.VideoList)

	c.JSON(http.StatusOK, response.FavoriteListResponse{
		Response: response.Response{
			StatusCode: 0,
		},
		VideoList: *favoriteList,
	})
}
