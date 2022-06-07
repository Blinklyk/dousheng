package model

import (
	"gorm.io/gorm"
	"time"
)

// Follow 关注表
type Follow struct {
	ID        int64          `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    int64          `json:"user_id"`
	FollowID  int64          `json:"follow_user_id""`
	Status    bool           `json:"status" gorm:"comment:1-follow each other"`
}
