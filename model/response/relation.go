package response

type FollowListResponse struct {
	Response
	UserList []UserInfo `json:"user_list"`
}
type FollowerListResponse struct {
	Response
	UserList []UserInfo `json:"user_list"`
}
