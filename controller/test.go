package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "conn success"})
}
