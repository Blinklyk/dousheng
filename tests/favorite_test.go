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

// test FavoriteList
func TestFavoriteListHandler(t *testing.T) {
	r := SetupRouter()
	//将项目中的API注册到测试使用的router
	r.GET("/favorite/list/", utils.JWTAuthMiddleware(), controller.FavoriteList)
	user_id := "2"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTY4LCJpc3MiOiJkb3V5aW4tZGVtbyIsImV4cCI6MTY1NTU3NDg2Mn0.uLNksH0U3pWp-kzqImVAiklWIbs2GhPNoaZpAluSFK8"
	//向注册的路由发起请求
	req, _ := http.NewRequest("GET", "/favorite/list/"+"?token="+token+"&user_id="+user_id, nil)
	w := httptest.NewRecorder()
	//模拟http服务处理请求
	global.App.DY_REDIS = initialize.InitializeRedis()
	log.Println(global.App.DY_REDIS)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

//func TestFavoriteActionHandler(t *testing.T) {
//	r := SetupRouter()
//	//将项目中的API注册到测试使用的router
//	r.POST("/favorite/action", utils.JWTAuthMiddleware(), controller.FavoriteAction)
//	video_id := "0"
//	action_type := "1"
//	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTY4LCJpc3MiOiJkb3V5aW4tZGVtbyIsImV4cCI6MTY1NTU3NDg2Mn0.uLNksH0U3pWp-kzqImVAiklWIbs2GhPNoaZpAluSFK8"
//	//向注册的路由发起请求
//	req, _ := http.NewRequest("POST", "/favorite/action/"+"?token="+token+"&video_id="+video_id+"&action_type="+action_type, nil)
//	w := httptest.NewRecorder()
//	global.App.DY_REDIS = initialize.InitializeRedis()
//	log.Println(global.App.DY_REDIS)
//	//模拟http服务处理请求
//	r.ServeHTTP(w, req)
//	assert.Equal(t, http.StatusOK, w.Code)
//}
