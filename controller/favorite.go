package controller

import (
	"encoding/json"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
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

	// call service
	fs := service.FavoriteService{}
	err := fs.FavoriteAction(&userInfoVar, &favoriteRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error in favorite action service: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{StatusCode: 0})
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

	// call service
	fs := service.FavoriteService{}
	favoriteVideoList, err := fs.FavoriteList(&favoriteListRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error: favoriteList service" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.FavoriteListResponse{
		Response: response.Response{
			StatusCode: 0,
		},
		VideoList: *favoriteVideoList,
	})
}
