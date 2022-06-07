package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/request"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"strconv"
)

type RelationService struct{}

// RelationAction two action type: add follow or cancel follow
func (us *RelationService) RelationAction(userInfoVar *model.User, actionVar *request.RelationActionRequest) error {

	actionType := actionVar.ActionType
	// follow operation
	if actionType == "1" {
		err := Follow(userInfoVar, actionVar)
		return err
	}

	if actionType == "2" {
		err := CancelFollow(userInfoVar, actionVar)
		return err
	}

	return nil

}

// Follow user adds follow relation with another user
func Follow(userInfoVar *model.User, actionVar *request.RelationActionRequest) error {
	toUserIDNum, _ := strconv.ParseInt(actionVar.ToUserID, 10, 64)

	// transaction: 1. if the table have the relation record; 2. determine the status column and set the toUser status 3. insert in follow table; 4. insert into follower table；
	// 5. user follow + 1 toUser follower + 1; 6. update userInfo in redis
	tx := global.DY_DB.Begin()

	// 1. if the table have the relation record;
	// check db first
	var followVar model.Follow
	res := tx.Model(&model.Follow{}).Where("user_id = ? AND follow_id = ?", userInfoVar.ID, actionVar.ToUserID).Find(&followVar)

	// have already followed
	if res.RowsAffected != 0 {
		return errors.New("error: have already followed")
	}

	// 2. determine the status column and set the toUser status
	var status bool
	// if in table relation, find the toUser follows the current user, status is true, otherwise false
	var toUserFollowVar model.Follow
	res = tx.Model(&model.Follow{}).Where("user_id = ? AND follow_id = ?", actionVar.ToUserID, userInfoVar.ID).First(&toUserFollowVar)
	if res.RowsAffected == 0 {
		status = false
	}
	if res.RowsAffected != 0 {
		status = true
		// set the toUser data status column to true (follow and follower table)
		if result := tx.Model(&model.Follow{}).Where("user_id = ? AND follow_id = ?", actionVar.ToUserID, userInfoVar.ID).Update("status", true); result.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("error: update status in toUser data in follow table")
		}

		if result := tx.Model(&model.Follower{}).Where("user_id = ? AND follower_id = ?", userInfoVar.ID, actionVar.ToUserID).Update("status", true); result.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("error: update status in toUser data in follower table")
		}
	}

	// 3. insert in follow table;
	newFollowVar := &model.Follow{
		UserID:   userInfoVar.ID,
		FollowID: toUserIDNum,
		Status:   status,
	}
	if res := tx.Model(&model.Follow{}).Create(newFollowVar); res.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("error: when create data in follow table")
	}

	// 4. insert into follower table

	newFollowerVar := &model.Follower{
		UserID:     toUserIDNum,
		FollowerID: userInfoVar.ID,
		Status:     status,
	}
	if res = tx.Model(&model.Follower{}).Create(newFollowerVar); res.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("error: when create data in follower table")
	}

	// 5. user follow + 1 toUser follower + 1;
	if res := tx.Model(&model.User{}).Where("id = ?", userInfoVar.ID).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)); res.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("error: when add follow_count in user table")
	}

	if res := tx.Model(&model.User{}).Where("id = ?", toUserIDNum).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)); res.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("error: when add follower_count in user table")
	}

	// 6. update userInfo in redis

	err := tx.Commit().Error
	return err
}

// CancelFollow user cancel follow relation with another user
func CancelFollow(userInfoVar *model.User, actionVar *request.RelationActionRequest) error {

	// transaction: 1. if the table have the relation record; 2. update the status column in follow table 3. delete relation in follow table; 4. delete relation in follower table;
	// 5. user follow + 1 toUser follower + 1; 6. update userInfo in redis

	// 1. if the table have the relation record;
	tx := global.DY_DB.Begin()

	// check db var
	var followVar model.Follow
	// check follow table has the corresponding follow record
	res := global.DY_DB.Model(&model.Follow{}).Where("user_id = ? AND follow_id = ?", userInfoVar.ID, actionVar.ToUserID).Find(&followVar)

	// if no relation data
	if res.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("did not follow before")
	}

	// 2. update status column of the toUser data in follow table
	// if toUser followed the current user, update status column from true to false in toUser follow record
	if followVar.Status == true {
		if res := tx.Model(&model.Follow{}).Where("user_id = ? AND follow_id = ?", actionVar.ToUserID, userInfoVar.ID).Update("status", false); res.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("error in updating the toUser status in follow table" + res.Error.Error())
		}
		if res := tx.Model(&model.Follower{}).Where("user_id = ? AND follower_id = ?", userInfoVar.ID, actionVar.ToUserID).Update("status", false); res.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("error in updating the toUser status in follower table" + res.Error.Error())
		}
	}

	// 3. delete relation in follow table
	if res := tx.Model(&model.Follow{}).Where("user_id = ? AND follow_id = ?", userInfoVar.ID, actionVar.ToUserID).Delete(&model.Follow{}); res.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("error in deleting record in follow table" + res.Error.Error())
	}

	// 4. delete relation in follower table;
	if res := tx.Model(&model.Follower{}).Where("user_id = ? AND follower_id = ?", actionVar.ToUserID, userInfoVar.ID).Delete(&model.Follower{}); res.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("error in deleting record in follower table: " + res.Error.Error())
	}

	// 5. user follow + 1 toUser follower + 1;
	if res := tx.Model(model.User{}).Where("id = ?", userInfoVar.ID).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		UpdateColumn("follow_count", gorm.Expr("follow_count + ?", -1)); res.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("error: when reduce follow_count in user table" + res.Error.Error())
	}

	if res := tx.Model(model.User{}).Where("id = ?", actionVar.ToUserID).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		UpdateColumn("follower_count", gorm.Expr("follower_count + ?", -1)); res.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("error: when reduce follower_count in user table" + res.Error.Error())
	}

	// 6. update userInfo in redis

	err := tx.Commit().Error
	return err
}

// FollowList get current user's follow list from db
func (us *RelationService) FollowList(r *request.FollowListRequest) (userList []model.User, err error) {
	// get follow user from follow table
	var followIDList []int64
	// query follow as a list in follow table
	if err := global.DY_DB.Table("dy_follow").Distinct("follow_id").Where("user_id = ? AND deleted_at is null", r.UserID).Find(&followIDList).Error; err != nil {
		return nil, err
	}
	log.Printf("%v\n", followIDList)
	// get follow user info from user table
	if err := global.DY_DB.Model(&model.User{}).Where("id in ?", followIDList).Find(&userList).Error; err != nil {
		return nil, err
	}

	// set all users is_follow to true
	for i := 0; i < len(userList); i++ {
		userList[i].IsFollow = true
	}
	return
}

// FollowerList get current user's follower list from db
func (us *RelationService) FollowerList(r *request.FollowerListRequest) (userList []model.User, err error) {
	var followerIDList []int64
	// query follower as a list in follow table
	if err := global.DY_DB.Table("dy_follower").Distinct("follower_id").Where("user_id = ? AND deleted_at is null", r.UserID).Find(&followerIDList).Error; err != nil {
		return nil, err
	}
	log.Printf("%v\n", followerIDList)
	// get follow user info from user table
	if err := global.DY_DB.Model(&model.User{}).Where("id in ?", followerIDList).Find(&userList).Error; err != nil {
		return nil, err
	}
	// traverse the user list
	for i := 0; i < len(userList); i++ {
		// check the status column in follower table
		var f model.Follower
		if err := global.DY_DB.Select("Status").Where("user_id = ? AND follower_id = ?", r.UserID, userList[i].ID).First(&f).Error; err != nil {
			return nil, err
		}
		if f.Status == true {
			userList[i].IsFollow = true
		}
	}
	return
}
