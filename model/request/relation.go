package request

type RelationActionRequest struct {
	Token      string `json:"token" form:"token"`
	ToUserID   string `json:"to_user_id" form:"to_user_id"`
	ActionType string `json:"action_type" form:"action_type"`
}

type FollowListRequest struct {
	Token  string `json:"token" form:"token"`
	UserID string `json:"user_id" form:"user_id"`
}

type FollowerListRequest struct {
	Token  string `json:"token" form:"token"`
	UserID string `json:"user_id" form:"user_id"`
}
