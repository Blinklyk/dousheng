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

	// call service
	fs := service.FavoriteService{}
	err := fs.FavoriteAction(&userInfoVar, &favoriteRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error in favorite action service: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{StatusCode: 0})
}

func GetVideoDTo(video model.Video) response.Video {
	var videoInfo response.Video
	videoInfo.ID = video.ID
	videoInfo.UserID = video.UserID
	videoInfo.User = GetUserDTo(video.User)
	videoInfo.PlayUrl = video.PlayUrl
	videoInfo.CoverUrl = video.CoverUrl
	videoInfo.FavoriteCount = video.FavoriteCount
	videoInfo.CommentCount = video.CommentCount
	videoInfo.IsFavorite = video.IsFavorite
	videoInfo.PublishTime = video.PublishTime
	videoInfo.Title = video.Title
	return videoInfo
}

func GetVideoListDTo(video []model.Video) []response.Video {
	videoInfo := make([]response.Video, len(video))
	for i := 0; i < len(video); i++ {
		videoInfo[i] = GetVideoDTo(video[i])
	}
	return videoInfo
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

	//verify
	if err := verify.IsNum(favoriteListRequest.UserID); err != nil {
		c.JSON(http.StatusBadRequest, Response{1, err.Error()})
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
		VideoList: GetVideoListDTo(*favoriteVideoList),
	})
}
