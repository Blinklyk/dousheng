package response

import "github.com/RaymondCode/simple-demo/model/dto"

type FollowListResponse struct {
	Response
	UserList []dto.UserDTO `json:"user_list"`
}
type FollowerListResponse struct {
	Response
	UserList []dto.UserDTO `json:"user_list"`
}
