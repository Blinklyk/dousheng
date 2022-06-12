package verify

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model/request"
	"regexp"
)

func oneOrTwo(num string) (err error) {
	if num == "1" || num == "2" {
		return nil
	}
	return fmt.Errorf("操作类型错误")
}
func IsNum(num string) (err error) {
	if result, _ := regexp.MatchString(`^[0-9]*$`, num); !result {
		return fmt.Errorf("应该全部都为数字")
	}
	return nil
}
func Email(emil string) (err error) {
	//校验用户账号（邮箱）
	if result, _ := regexp.MatchString(`^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, emil); !result {
		return fmt.Errorf("Please enter the correct mailbox")
	}
	if result, _ := regexp.MatchString(`^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, emil); !result {
		return fmt.Errorf("Please enter the correct mailbox")
	}
	return nil
}

//密码强度必须为字⺟⼤⼩写+数字+符号，8位以上
func PassWord(Password string) (err error) {
	//校验密码
	if len(Password) < 8 {
		return fmt.Errorf("password len is < 9")
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	if b, err := regexp.MatchString(num, Password); !b || err != nil {
		return fmt.Errorf("password need number")
	}
	if b, err := regexp.MatchString(a_z, Password); !b || err != nil {
		return fmt.Errorf("password need a_z ")
	}
	if b, err := regexp.MatchString(A_Z, Password); !b || err != nil {
		return fmt.Errorf("password need A_Z")
	}
	return nil
}
func null(str string) (err error) {
	if reault, _ := regexp.MatchString(`^\s*$`, str); reault {
		return fmt.Errorf("内容不能为空")
	}
	return nil
}
func Resgin(reg request.RegisterRequest) (err error) {
	if err := Email(reg.Username); err != nil {
		return err
	}
	if err = PassWord(reg.Password); err != nil {
		return err
	}
	return nil
}

func Login(login request.LoginRequest) (err error) {
	if err := Email(login.Username); err != nil {
		return err
	}
	if err = PassWord(login.Password); err != nil {
		return err
	}
	return nil
}
func Comment(commentRequest request.CommentRequest) (err error) {
	if err := IsNum(commentRequest.VideoID); err != nil {
		return fmt.Errorf("视频ID错误")
	}
	if err := oneOrTwo(commentRequest.ActionType); err != nil {
		return err
	}
	if err := null(commentRequest.CommentText); err != nil {
		return fmt.Errorf("内容不能为空")
	}
	return
}
func Favorite(favoriteRequest request.FavoriteRequest) (err error) {
	if err := IsNum(favoriteRequest.VideoID); err != nil {
		return err
	}
	if err := oneOrTwo(favoriteRequest.ActionType); err != nil {
		return nil
	}
	return
}
func Publish(publishRequest request.PublishRequest) (err error) {
	if publishRequest.Title == "" {
		return fmt.Errorf("标题不能为空")
	}
	return nil
}
func Relation(request request.RelationActionRequest) (err error) {
	if err := IsNum(request.ToUserID); err != nil {
		return fmt.Errorf("对方用户不合法")
	}
	if err := oneOrTwo(request.ActionType); err != nil {
		return fmt.Errorf("操作类型错误")
	}
	return
}
func User() (err error) {
	return
}
