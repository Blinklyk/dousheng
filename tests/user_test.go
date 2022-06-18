package tests

import (
	"bytes"
	"encoding/json"
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/initialize"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// test commentAction
func TestRegisterHandler(t *testing.T) {
	r := SetupRouter()
	//将项目中的API注册到测试使用的router
	r.POST("/user/register", controller.Register)
	registerReq := request.LoginRequest{
		Username: "fefe13212321123",
		Password: "Lzj2322193127",
	}
	//序列化请求体
	jsonValue, _ := json.Marshal(registerReq)
	//向注册的路由发起请求
	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Accept", "*/*")
	// mandatory here
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	global.App.DY_REDIS = initialize.InitializeRedis()
	log.Println(global.App.DY_REDIS)
	//模拟http服务处理请求
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLoginHandler(t *testing.T) {
	r := SetupRouter()
	//将项目中的API注册到测试使用的router
	r.POST("/user/login/", controller.Login)
	loginReq := request.LoginRequest{
		Username: "fefe",
		Password: "Lzj2322193127",
	}
	//序列化请求体
	jsonValue, _ := json.Marshal(loginReq)
	//向注册的路由发起请求
	req, _ := http.NewRequest("POST", "/user/login/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	global.App.DY_REDIS = initialize.InitializeRedis()
	log.Println(global.App.DY_REDIS)
	//模拟http服务处理请求
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserInfoHandler(t *testing.T) {
	r := SetupRouter()
	//将项目中的API注册到测试使用的router
	r.GET("/user/", controller.Login)
	user_id := "1"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTY4LCJpc3MiOiJkb3V5aW4tZGVtbyIsImV4cCI6MTY1NTU3MTIzN30.3FKu64sjamB9wo9LP604HlEHpnU6PkjjV7myneJO9DI"
	//向注册的路由发起请求
	req, _ := http.NewRequest("GET", "/user/"+"?user_id="+user_id+"&token="+token, nil)
	w := httptest.NewRecorder()
	global.App.DY_REDIS = initialize.InitializeRedis()
	log.Println(global.App.DY_REDIS)
	//模拟http服务处理请求
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
