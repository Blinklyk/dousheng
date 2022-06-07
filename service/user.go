package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/utils"
	"gorm.io/gorm"
	"log"
)

type UserService struct{}

// Register user register and store to db
func (us *UserService) Register(user *model.User) (err error, newUser *model.User) {
	// 校验 查询数据库中是否有此用户(高级查询)
	var u model.User
	if !errors.Is(global.DY_DB.Model(&model.User{}).Where("username = ?", user.Username).First(&u).Error, gorm.ErrRecordNotFound) {
		return errors.New("this username is registered already"), user
	}
	// 雪花算法生成新的id
	//var node, _ = utils.NewWorker(1)
	//newID := node.GetId()
	//user.ID = newID
	// 密码加密
	user.Password = utils.BcryptHash(user.Password)
	// 添加到数据库
	log.Printf("%v\n", user)
	err = global.DY_DB.Create(&user).Error
	return err, user
}

// Login user login and store some date to redis
func (us *UserService) Login(user *model.User) (returnUser *model.User, tokenStr string, err error) {

	// jwt version
	// TODO 校验
	// 查询 账号密码是否正确
	var u model.User

	// get user form db
	if errors.Is(global.DY_DB.Model(&model.User{}).Where("username = ?", user.Username).First(&u).Error, gorm.ErrRecordNotFound) {
		return nil, "", errors.New("user doesn't exist")
	}
	if ok := utils.BcryptCheck(user.Password, u.Password); !ok {
		return nil, "", errors.New("password error")
	}
	// forbid the FollowCount/FollowerCount get into redis with session info
	u.FollowerCount = -1
	u.FollowCount = -1
	log.Printf("get User data from db : %v", u)
	// gen token
	tokenStr, err = utils.GenToken(u.ID)
	if err != nil {
		return nil, "", err
	}

	// store user data into redis
	jsonU, err := json.Marshal(u)
	if err != nil {
		return nil, "", errors.New("json marshal error")
	}
	// redis key: "login:session:"+tokenStr, value: user TTL: 30min
	res := global.DY_REDIS.Set(context.Background(), global.REDIS_USER_PREFIX+tokenStr, jsonU, global.REDIS_USER_TTL)
	log.Println("res.String() user set to redis:", res)
	return &u, tokenStr, nil

	//// session + redis version
	//// TODO check format
	//var u model.User
	//if errors.Is(global.DY_DB.Model(&model.User{}).Where("username = ?", user.Username).First(&u).Error, gorm.ErrRecordNotFound) {
	//	return nil, "", errors.New("user doesn't exist")
	//}
	//if ok := utils.BcryptCheck(user.Password, u.Password); !ok {
	//	return nil, "", errors.New("password error")
	//}
	//// 生成session key : userID, value: user
	//jsonU, err := json.Marshal(u)
	//if err != nil {
	//	return nil, "", errors.New("json marshal error")
	//}
	//session.Set(u.ID, jsonU)
	//err = session.Save()
	//if err != nil {
	//	return nil, "", err
	//}
	//return &u, strconv.FormatInt(u.ID, 10), nil

}

// GetUserInfo get full latest userInfo from db (follow_count + follower_count + is_followed)
func (us *UserService) GetUserInfo(userID int64, toUserID int64) (returnUser *model.User, err error) {
	var u model.User
	// get user basic info from user table
	if errors.Is(global.DY_DB.Model(&model.User{}).Where("id = ?", userID).First(&u).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user doesn't exist")
	}
	// get user "is_follow" column from follow table
	if res := global.DY_DB.Model(&model.Follow{}).Where("user_id = ? AND follow_id = ?", userID, toUserID).First(&u); res.RowsAffected == 0 {
		u.IsFollow = false
	} else {
		u.IsFollow = true
	}

	return &u, nil

}
