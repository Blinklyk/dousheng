package request

type FavoriteRequest struct {
	Token      string `json:"token" form:"token"`
	VideoID    string `json:"video_id" form:"video_id"`
	ActionType string `json:"action_type" form:"action_type"`
}

type FavoriteListRequest struct {
	Token  string `json:"token" form:"token"`
	UserID int64  `json:"user_id" form:"user_id"`
}
