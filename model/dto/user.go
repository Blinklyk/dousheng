package dto

type UserDTO struct {
	ID            int64  `json:"id"`
	Name          string `json:"name" gorm:"default:testName"`
	FollowCount   int64  `json:"follow_count" gorm:"default:0"`   // 关注数
	FollowerCount int64  `json:"follower_count" gorm:"default:0"` // 粉丝数
	IsFollow      bool   `json:"is_follow" gorm:"default:false"`  // 当前用户是否关注
}
