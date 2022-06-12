package model

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	// TODO 改成嵌入结构体
	ID            int64          `gorm:"primarykey"` // 主键ID
	CreatedAt     time.Time      // 创建时间
	UpdatedAt     time.Time      // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`                                      // 删除时间
	UserID        int64          `json:"author_id,omitempty"`                                 // 发布作者
	User          User           `json:"user,omitempty" gorm:"foreignKey:UserID"`             // user信息
	PlayUrl       string         `json:"play_url,omitempty" gorm:"default_playUrl"`           // 视频地址
	CoverUrl      string         `json:"cover_url,omitempty" gorm:"default:default_coverUrl"` // 封面地址
	FavoriteCount int64          `json:"favorite_count,omitempty" gorm:"default:0"`           // 点赞数量
	CommentCount  int64          `json:"comment_count,omitempty" gorm:"default:0"`            // 评论数量
	IsFavorite    bool           `json:"is_favorite,omitempty" gorm:"default:false"`          // 是否点赞
	PublishTime   time.Time      `json:"publish_time,omitempty" gorm:"comment:投稿时间"`          // 投稿时间
	Title         string         `json:"title, omitempty" gorm:"comment:视频说明"`                // 投稿时添加的文字
	CommentList   []Comment      `json:"comment_list"`                                        // 视频下的评论列表
}
