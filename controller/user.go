package controller

import (
	"encoding/json"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {

	// bind request var
	var r request.RegisterRequest
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call service
	newUser := &model.User{Username: r.Username, Password: r.Password, FollowCount: 0, FollowerCount: 0}
	var userService = service.UserService{}
	err, userReturn := userService.Register(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.RegisterResponse{Response: response.Response{StatusCode: -1, StatusMsg: "failed: create register user " + err.Error()}})
		return
	}

	// return
	token, _ := utils.GenToken(userReturn.ID)
	c.JSON(http.StatusOK, response.RegisterResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "success: create register user",
		},
		UserId: userReturn.ID,
		// TODO register token
		Token: token,
	})
}

func Login(c *gin.Context) {

	// bind request var
	var l request.LoginRequest
	if err := c.ShouldBind(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call service
	user := &model.User{Username: l.Username, Password: l.Password}
	var loginService = service.UserService{}
	userReturn, tokenStr, err := loginService.Login(user)
	if tokenStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error tokenStr is empty": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, response.LoginResponse{Response: response.Response{StatusCode: -1, StatusMsg: "failed: login in" + err.Error()}})
		return
	}

	// return
	c.JSON(http.StatusOK, response.LoginResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "success: login in",
		},
		UserId: userReturn.ID,
		Token:  tokenStr,
	})

	return
}

// UserInfo get login userInfo from db
func UserInfo(c *gin.Context) {

	// authentication jwt version
	UserStr, _ := c.Get("UserStr")

	var userInfoVar model.User
	if err := json.Unmarshal([]byte(UserStr.(string)), &userInfoVar); err != nil {
		c.JSON(http.StatusOK, response.UserInfoResponse{Response: response.Response{StatusCode: 1, StatusMsg: "error: session unmarshal error"}})
		return
	}

	// TODO check
	if len(userInfoVar.Name) < 3 {
		c.JSON(http.StatusOK, response.UserInfoResponse{
			Response: response.Response{StatusCode: 1, StatusMsg: "userName len less then 3"},
		})
		return
	}

	// call service
	var checkUserInfoService = service.UserService{}
	returnUser, err := checkUserInfoService.GetUserInfo(userInfoVar.ID, userInfoVar.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 1, StatusMsg: "error: db select"})
		return
	}
	//type UserResp struct {
	//	response.Response
	//	// TODO 固定字段改成嵌入结构体
	//	ID             int64         `json:"id"`
	//	Name           string        `json:"name,omitempty" gorm:"default:testName"`    // TODO
	//	FollowCount    int64         `json:"follow_count,omitempty" gorm:"default:0"`   // 关注数
	//	FollowerCount  int64         `json:"follower_count,omitempty" gorm:"default:0"` // 粉丝数
	//	IsFollow       bool          `json:"is_follow,omitempty" gorm:"default:false"`  // 当前用户是否关注
	//	Username       string        `json:"username" gorm:"comment:username" `         // 登录账号
	//	Videos         []model.Video `json:"videos"`                                    // 发布视频列表
	//	FavoriteVideos []model.Video `json:"favorite_videos"`                           //`gorm:"many2many:favorite"`
	//}
	//c.JSON(http.StatusOK, UserResp{
	//	Response:       response.Response{StatusCode: 0},
	//	ID:             returnUser.ID,
	//	Name:           returnUser.Name,
	//	FollowCount:    returnUser.FollowCount,
	//	FollowerCount:  returnUser.FollowerCount,
	//	IsFollow:       returnUser.IsFollow,
	//	Username:       returnUser.Username,
	//	Videos:         returnUser.Videos,
	//	FavoriteVideos: returnUser.FavoriteVideos,
	//})
	userinfo := response.UserInfo{
		ID:             returnUser.ID,
		Name:           returnUser.Name,
		FollowCount:    returnUser.FollowCount,
		FollowerCount:  returnUser.FollowerCount,
		IsFollow:       returnUser.IsFollow,
		Username:       returnUser.Username,
		Videos:         returnUser.Videos,
		FavoriteVideos: returnUser.FavoriteVideos,
	}
	c.JSON(http.StatusOK, response.UserInfoResponse{
		Response: response.Response{StatusCode: 0},
		UserInfo: userinfo,
	})
	return

	////session version
	//// get user from redis
	//userID := c.Query("user_id")
	//session := sessions.Default(c)
	//jsonUser := session.Get(userID)
	//log.Println("jsonUser : ", jsonUser)
	//log.Println(c.GetHeader("Cookie"))
	//log.Println(c.GetHeader("Host"))
	//log.Println(c.GetHeader("Connection"))
	//
	////userInfoVar := &userInfoVar{}
	//userInfoVar := &model.User{}
	//
	//err := json.Unmarshal(jsonUser.([]byte), userInfoVar)
	//if err != nil {
	//	c.JSON(http.StatusOK, CheckUserInfoResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "Unmarshal from session failed"},
	//	})
	//	return
	//}
	//
	//if len(userInfoVar.Name) < 0 {
	//	c.JSON(http.StatusOK, CheckUserInfoResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "userName len is 0"},
	//	})
	//	return
	//}
	//
	//
	//c.JSON(http.StatusOK, CheckUserInfoResponse{
	//	Response: Response{StatusCode: 0},
	//	UserInfo: *userInfoVar,
	//})
	//return

}
