package controller

import (
	"encoding/json"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils/verify"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// RelationAction actionType = 1: follow; actionType = 2: cancel follow
func RelationAction(c *gin.Context) {
	UserStr, _ := c.Get("UserStr")
	log.Println("UserStr: ", UserStr)

	var userInfoVar model.User
	if err := json.Unmarshal([]byte(UserStr.(string)), &userInfoVar); err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "error: session unmarshal error"})
		return
	}

	var relationActionRequest request.RelationActionRequest
	if err := c.ShouldBind(&relationActionRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "bind error"})
		return
	}

	if err := verify.Relation(relationActionRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{1, err.Error()})
		return
	}
	if err := c.ShouldBind(&relationActionRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "bind error"})
		return
	}
	log.Printf("%v\n", relationActionRequest)

	// verify

	// cannot follow myself
	if relationActionRequest.ToUserID == strconv.Itoa(int(userInfoVar.ID)) {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "cannot follow myself"})
		return
	}

	relationService := &service.RelationService{}
	if err := relationService.RelationAction(&userInfoVar, &relationActionRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0})

}

// FollowList get follow list of the current user
func FollowList(c *gin.Context) {
	UserStr, _ := c.Get("UserStr")
	log.Println("UserStr: ", UserStr)

	var userInfoVar model.User
	if err := json.Unmarshal([]byte(UserStr.(string)), &userInfoVar); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error: session unmarshal error"})
		return
	}

	var followListRequest request.FollowListRequest
	if err := c.ShouldBind(&followListRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "bind error"})
		return
	}

	relationService := &service.RelationService{}
	followList, err := relationService.FollowList(&followListRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.FollowListResponse{
		Response: response.Response{
			StatusCode: 0,
		},
		UserList: followList,
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

	var followerListRequest request.FollowerListRequest
	if err := c.ShouldBind(&followerListRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "bind error"})
		return
	}

	relationService := &service.RelationService{}
	followerList, err := relationService.FollowerList(&followerListRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: err.Error()})
	}

	c.JSON(http.StatusOK, response.FollowListResponse{
		Response: response.Response{
			StatusCode: 0,
		},
		UserList: followerList,
	})

}
