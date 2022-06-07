package request

type CommentRequest struct {
	Token       string `json:"token" form:"token"`
	VideoID     string `json:"video_id" form:"video_id"`
	ActionType  string `json:"action_type" form:"action_type"`
	CommentText string `json:"comment_text" form:"comment_text"`
	CommentID   string `json:"comment_id" form:"comment_id"`
}

type CommentListRequest struct {
	Token   string `json:"token" form:"token"`
	VideoID string `json:"video_id" form:"video_id"`
}
