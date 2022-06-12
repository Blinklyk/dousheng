package response

import (
	"github.com/RaymondCode/simple-demo/model"
	"time"
)

type VideoDTO struct {
	// TODO 改成嵌入结构体
	ID            int64           `gorm:"primarykey"`                                  // 主键ID
	UserID        int64           `json:"author_id,omitempty"`                         // 发布作者
	User          UserInfo        `json:"author,omitempty" gorm:"foreignKey:UserID"`   // user信息
	PlayUrl       string          `json:"play_url,omitempty" gorm:"default:testName"`  // 视频地址
	CoverUrl      string          `json:"cover_url,omitempty" gorm:"default:testName"` // 封面地址
	FavoriteCount int64           `json:"favorite_count" gorm:"default:0"`             // 点赞数量
	CommentCount  int64           `json:"comment_count" gorm:"default:0"`              // 评论数量
	IsFavorite    bool            `json:"is_favorite" gorm:"default:false"`            // 是否点赞
	PublishTime   time.Time       `json:"publish_time" gorm:"comment:投稿时间"`            // 投稿时间
	Title         string          `json:"title, omitempty" gorm:"comment:视频说明"`        // 投稿时添加的文字
	CommentList   []model.Comment `json:"comment_list,omitempty"`                      // 视频下的评论列表
}
