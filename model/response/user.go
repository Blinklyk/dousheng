package response

import "github.com/RaymondCode/simple-demo/model"

type RegisterResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type LoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfo struct {
	// TODO 固定字段改成嵌入结构体
	ID             int64         `json:"id"`
	Name           string        `json:"name,omitempty" gorm:"default:testName"`    // TODO
	FollowCount    int64         `json:"follow_count,omitempty" gorm:"default:0"`   // 关注数
	FollowerCount  int64         `json:"follower_count,omitempty" gorm:"default:0"` // 粉丝数
	IsFollow       bool          `json:"is_follow,omitempty" gorm:"default:false"`  // 当前用户是否关注
	Username       string        `json:"username" gorm:"comment:username" `         // 登录账号
	Videos         []model.Video `json:"videos"`                                    // 发布视频列表
	FavoriteVideos []model.Video `json:"favorite_videos"`                           //`gorm:"many2many:favorite"`
}
type UserInfoResponse struct {
	Response
	UserInfo UserInfo `json:"user_info"`
}
