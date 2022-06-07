package request

type PublishRequest struct {
	Token string `json:"token" form:"token"`
	Title string `json:"title" form:"title"`
}

type PublishListRequest struct {
	Token  string `json:"token" form:"token"`
	UserID string `json:"user_id" form:"user_id"`
}
