package model

import (
	"gorm.io/gorm"
	"time"
)

type Favorite struct {
	ID        int64 `gorm:"primaryKey"`
	UserID    int64
	VideoID   int64
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
