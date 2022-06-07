package controller

import (
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
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

	// call service
	fs := service.FeedService{}
	token := feedRequest.Token
	// if request doesn't contain token
	if token == "" {
		feedList, err := fs.FeedWithoutToken()
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error:feed without token" + err.Error()})
			return
		}
		c.JSON(http.StatusOK, response.FeedResponse{
			Response:  response.Response{StatusCode: 0},
			VideoList: *feedList,
			NextTime:  time.Now().Unix(),
		})
		return
	}

	// if request contains token
	if token != "" {
		feedList, err := fs.FeedWithToken(&feedRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error:feed with token" + err.Error()})
			return
		}
		c.JSON(http.StatusOK, response.FeedResponse{
			Response:  response.Response{StatusCode: 0},
			VideoList: *feedList,
			NextTime:  time.Now().Unix(),
		})
		return
	}

}
