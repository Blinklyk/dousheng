package rpcdto

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/pb/rpcUser"
)

// ToUserRpcDTO model user to rpc user
func ToUserRpcDTO(user *model.User) *rpcUser.User {
	userReturn := &rpcUser.User{
		Id:            user.ID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
		Username:      user.Username,
		Password:      user.Password,
	}
	return userReturn
}

func ToUserListDTO(users *[]model.User) []*rpcUser.User {
	usersReturn := make([]*rpcUser.User, len(*users), len(*users))
	// traverse the user list and add v to usersReturn list
	for i := 0; i < len(*users); i++ {
		user := *users
		u := &rpcUser.User{
			Id:            user[i].ID,
			Name:          user[i].Name,
			FollowCount:   user[i].FollowCount,
			FollowerCount: user[i].FollowerCount,
			IsFollow:      user[i].IsFollow,
			Username:      user[i].Username,
			Password:      user[i].Password,
		}
		usersReturn[i] = u

	}
	return usersReturn
}
