package tests

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/initialize"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPublishListHandler(t *testing.T) {
	r := SetupRouter()
	//将项目中的API注册到测试使用的router
	r.GET("/publish/list/", utils.JWTAuthMiddleware(), controller.PublishList)
	user_id := "1"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTY4LCJpc3MiOiJkb3V5aW4tZGVtbyIsImV4cCI6MTY1NTU3NDg2Mn0.uLNksH0U3pWp-kzqImVAiklWIbs2GhPNoaZpAluSFK8"
	//向注册的路由发起请求
	req, _ := http.NewRequest("GET", "/publish/list/"+"?user_id="+user_id+"&token="+token, nil)
	w := httptest.NewRecorder()
	global.App.DY_REDIS = initialize.InitializeRedis()
	log.Println(global.App.DY_REDIS)
	//模拟http服务处理请求
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
