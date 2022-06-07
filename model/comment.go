package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID         int64          `gorm:"primarykey"` // 主键ID
	CreatedAt  time.Time      // 创建时间
	UpdatedAt  time.Time      // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	UserID     int64          `json:"user_id" `
	User       User           `json:"user" gorm:"foreignKey:UserID"`
	VideoID    int64          `json:"video_id"`
	Content    string         `json:"content"`
	CreateData string         `json:"create_data"`
}
