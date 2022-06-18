package tests

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// test code

//为测试使用创建 *gin.Engine实例
func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}
func TestFeedHandler(t *testing.T) {
	r := SetupRouter()
	//将项目中的API注册到测试使用的router
	r.GET("/feed/", controller.Feed)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjIsImlzcyI6ImRvdXlpbi1kZW1vIiwiZXhwIjoxNjU1NTQ4MTgzfQ.93Xd9CxFxgWF91VSYwLYWUzcyT-vH01loabiTYajsX0"
	//向注册的路由发起请求
	req, _ := http.NewRequest("GET", "/feed/"+"?token="+token, nil)
	w := httptest.NewRecorder()

	//模拟http服务处理请求
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
