package response

type CommentInfo struct {
	ID         int64    `gorm:"primarykey"` // 主键ID
	UserID     int64    `json:"user_id" `
	User       UserInfo `json:"user" gorm:"foreignKey:UserID"`
	VideoID    int64    `json:"video_id"`
	Content    string   `json:"content"`
	CreateData string   `json:"create_data"`
}

type CommentActionResponse struct {
	Response
	Comment CommentInfo `json:"comment"`
}

type CommentListResponse struct {
	Response
	CommentList []CommentInfo `json:"comment_list"`
}
