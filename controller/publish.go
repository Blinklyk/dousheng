package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/bwmarrin/snowflake"

	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils/verify"
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
=======
	"go.uber.org/zap"
	"net/http"
	"path/filepath"
>>>>>>> a5ad9421cddcb4c71a3ebda7d6ed77f835c4b828
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	// authentication
	UserStr, _ := c.Get("UserStr")

	var userInfoVar model.User
	if err := json.Unmarshal([]byte(UserStr.(string)), &userInfoVar); err != nil {
		global.App.DY_LOG.Error("session unmarshal error", zap.Error(err))
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "error: session unmarshal error"})
		return
	}

	// bind request var
	var publicRequest request.PublishRequest
	if err := c.ShouldBind(&publicRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "bind error " + err.Error()})
		return
	}

	//verify
	if err := verify.Publish(publicRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{1, "非法数据"})
		return
	}

	// save the file at localhost and get the localFilePath
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "error: data " + err.Error()})
		return
	}
	filename := filepath.Base(data.Filename)
	indexOfDot := strings.LastIndex(filename, ".") // get suffix index
	if indexOfDot < 0 {
		log.Println("error: can't get suffix")
		return
	}
	suffix := filename[indexOfDot+1 : len(filename)]
	suffix = strings.ToLower(suffix)

	// generate newfilename
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	filename = strconv.FormatInt(node.Generate().Int64(), 10)
	filename = filename + "." + suffix

	log.Println("filename: ", filename)

	localFilePath := filepath.Join(global.LOCAL_FILE_PATH_PREFIX, filename)
	if err := c.SaveUploadedFile(data, localFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error: save upload file " + err.Error()})
		return
	}

	// call publish action service
	ps := service.PublishService{}
	if err = ps.PublishAction(&userInfoVar, &publicRequest, localFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error in publish service: " + err.Error()})
	}

	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "发布成功！"})
	return

}

// Get PublishList

func PublishList(c *gin.Context) {

	// authentication
	UserStr, _ := c.Get("UserStr")

	var userInfoVar model.User
	if err := json.Unmarshal([]byte(UserStr.(string)), &userInfoVar); err != nil {
		global.App.DY_LOG.Error("session unmarshal error", zap.Error(err))
		c.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: "error: session unmarshal error"})
		return
	}

	// bind request var
	var publishListRequest request.PublishListRequest
	if err := c.ShouldBind(&publishListRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "bind error " + err.Error()})
		return
	}

	if err := verify.IsNum(publishListRequest.UserID); err != nil {
		c.JSON(http.StatusBadRequest, Response{1, err.Error()})
		return
	}

	// call service
	ps := service.PublishService{}
	publishVideos, err := ps.PublishList(&publishListRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error in publish service: " + err.Error()})
		return
	}

	videosInfo := GetVideoListDTo(publishVideos)
	// return
	c.JSON(http.StatusOK, response.PublishListResponse{
		Response: response.Response{
			StatusCode: 0,
		},
		VideoList: videosInfo,
	})
}
