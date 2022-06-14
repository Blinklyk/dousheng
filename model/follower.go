package model

import (
	"gorm.io/gorm"
	"time"
)

// Follower 粉丝表
type Follower struct {
	ID         int64          `gorm:"primarykey"` // 主键ID
	CreatedAt  time.Time      // 创建时间
	UpdatedAt  time.Time      // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	UserID     int64          `json:"user_id"`
	FollowerID int64          `json:"follower_user_id""`
	Status     bool           `json:"status" gorm:"comment:1-follow each other"`
}
