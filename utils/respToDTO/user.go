package respToDTO

import (
	"github.com/RaymondCode/simple-demo/model/dto"
	"github.com/RaymondCode/simple-demo/pb/rpcUser"
)

func GetUserDTo(user *rpcUser.User) (userInfo *dto.UserDTO) {
	userInfo = &dto.UserDTO{
		ID:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
	return userInfo
}

func GetUserListDTO(users []*rpcUser.User) *[]dto.UserDTO {
	userList := make([]dto.UserDTO, len(users))
	for i := 0; i < len(users); i++ {
		userList[i] = *GetUserDTo(users[i])
	}
	return &userList
}
