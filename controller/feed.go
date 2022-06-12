package controller

import (
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Feed token is optional here
func Feed(c *gin.Context) {

	// bind request var
	var feedRequest request.FeedRequest
	if err := c.ShouldBind(&feedRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "feed bind error"})
	}
	log.Printf("feedRequest%v\n", feedRequest)
	latestTime := c.Query("latest_time")
	log.Println("latest_time: ", latestTime)

	// call service
	fs := service.FeedService{}
	token := c.Query("token")
	// if request doesn't contain token
	if token == "" {
		feedList, err := fs.FeedWithoutToken()
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error:feed without token" + err.Error()})
			return
		}
		videoInfo := GetVideoListDTo(*feedList)
		log.Println(" withToken nextTime:", videoInfo[len(videoInfo)-1].PublishTime.Unix(), "   ", videoInfo[len(videoInfo)-1].PublishTime)
		c.JSON(http.StatusOK, response.FeedResponse{
			Response:  response.Response{StatusCode: 0},
			VideoList: videoInfo,
			NextTime:  videoInfo[len(videoInfo)-1].PublishTime.Unix(),
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
		videoInfo := GetVideoListDTo(*feedList)
		log.Println(" withToken nextTime:", videoInfo[len(videoInfo)-1].PublishTime.Unix(), videoInfo[len(videoInfo)-1].PublishTime)
		c.JSON(http.StatusOK, response.FeedResponse{
			Response:  response.Response{StatusCode: 0},
			VideoList: videoInfo,
			NextTime:  videoInfo[len(videoInfo)-1].PublishTime.Unix(),
		})
	}
}
