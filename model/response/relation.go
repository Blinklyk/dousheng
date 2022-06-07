package response

import "github.com/RaymondCode/simple-demo/model"

type FollowListResponse struct {
	Response
	UserList []model.User `json:"user_list"`
}

type FollowerListResponse struct {
	Response
	UserList []model.User `json:"user_list"`
}
