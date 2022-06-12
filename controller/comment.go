package controller

import (
	"encoding/json"
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils/verify"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 1, StatusMsg: "bind error "})
		return
	}

	//verify
	if err := verify.Comment(commentRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 1, StatusMsg: "非法数据 "})
		return
	}

	// call service: action_type = 1 add comment; action_type = 2 delete comment
	cs := service.CommentService{}
	actionType := commentRequest.ActionType
	if actionType == "1" {
		commentVar, err := cs.CommentAction(&userInfoVar, &commentRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error in commentAction: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, response.CommentActionResponse{
			Response: response.Response{StatusCode: 0},
			Comment:  GetCommentDTo(*commentVar),
		})
		return
	}

	// delete comment
	if actionType == "2" {
		if err := cs.DeleteCommentAction(&commentRequest); err != nil {
			c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error in deleteCommentAction: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
		return
	}

}

func GetCommentDTo(comment model.Comment) (commentInfo response.CommentInfo) {
	commentInfo.ID = comment.ID
	commentInfo.UserID = comment.UserID
	commentInfo.User = GetUserDTo(comment.User)
	commentInfo.VideoID = comment.VideoID
	commentInfo.Content = comment.Content
	commentInfo.CreateData = comment.CreateData
	return
}

func GetCommentListInfo(commentList *[]model.Comment) []response.CommentInfo {
	//var commentListInfo [length]response.CommentInfo
	commentListInfo := make([]response.CommentInfo, len(*commentList))
	//commentListInfo := make([]response.CommentInfo, 1, length)
	fmt.Println(*commentList)
	for index, value := range *commentList {
		//commentListInfo[index] = GetCommentDTo(value)
		commentListInfo[index].ID = value.ID
		commentListInfo[index].UserID = value.UserID
		commentListInfo[index].User = GetUserDTo(value.User)
		commentListInfo[index].VideoID = value.VideoID
		commentListInfo[index].Content = value.Content
		commentListInfo[index].CreateData = value.CreateData
	}
	return commentListInfo
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

	//verify
	if err := verify.IsNum(commentListRequest.VideoID); err != nil {
		c.JSON(http.StatusBadRequest, Response{1, err.Error()})
		return
	}

	// call service
	cs := service.CommentService{}
	commentList, err := cs.CommentList(userInfoVar, &commentListRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error in commentList: " + err.Error()})
	}
	// return
	c.JSON(http.StatusOK, response.CommentListResponse{
		Response:    response.Response{StatusCode: 0},
		CommentList: GetCommentListInfo(commentList),
	})
}
