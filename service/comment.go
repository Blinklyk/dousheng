package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/request"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
)

type CommentService struct{}

// CommentAction add comment return new comment
func (cs *CommentService) CommentAction(userId int64, r *request.CommentRequest) (*model.Comment, error) {
	videoID := r.VideoID
	commentText := r.CommentText
	videoIDNum, _ := strconv.ParseInt(videoID, 10, 64)
	// 1. create data in comment table; 2. corresponding video comment_count + 1
	tx := global.App.DY_DB.Begin()

	// get full user data
	userService := UserService{}
	returnUser, err := userService.GetUserInfo(userId, userId)
	if err != nil {
		return nil, errors.New("error: return user from getUserInfo")
	}
	commentVar := &model.Comment{
		UserID:     userId,
		VideoID:    videoIDNum,
		Content:    commentText,
		User:       *returnUser,
		CreateData: time.Now().Format("01-02"),
	}
	// 1. create data in comment table
	if res := tx.Model(&model.Comment{}).Create(&commentVar); res.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("error: create comment in comment table")
	}

	// 2. corresponding video comment_count + 1
	// update the comment_count column (lock) in video table
	if res := tx.Model(&model.Video{}).Where("id = ?", videoID).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)); res.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("\"error: update add comment_count in video table\"")
	}
	if err := tx.Commit().Error; err != nil {
		return nil, errors.New("error: commit update transaction: " + err.Error())
	}
	return commentVar, nil
}

// DeleteCommentAction delete comment
func (cs *CommentService) DeleteCommentAction(r *request.CommentRequest) error {
	videoID := r.VideoID
	tx := global.App.DY_DB.Begin()
	commentID := r.CommentID
	// 1. delete comment in comment table
	if res := tx.Delete(&model.Comment{}, "id = ? AND video_id = ?", commentID, videoID); res.RowsAffected == 0 {
		return errors.New("err: didn't get this comment")
	}
	// 2. update comment_count - 1 in video table
	if res := tx.Model(&model.Video{}).Where("id = ?", videoID).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", -1)); res.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("error: update subtract comment_count in video table")
	}
	if err := tx.Commit().Error; err != nil {
		return errors.New("error: commit delete transaction")
	}
	return nil
}

// CommentList get comment list
func (cs *CommentService) CommentList(userId int64, r *request.CommentListRequest) (*[]model.Comment, error) {
	videoID := r.VideoID
	var commentList []model.Comment
	if err := global.App.DY_DB.Model(&model.Comment{}).Where("video_id = ?", videoID).Preload("User").Find(&commentList).Error; err != nil {
		return nil, errors.New("get comment list when db select error")
	}

	// add is_follow value to the comment list
	for i := 0; i < len(commentList); i++ {
		var temp model.Follow
		if res := global.App.DY_DB.Model(&model.Follow{}).Where("user_id = ? AND follow_id = ?", userId, commentList[i].UserID).First(&temp); res.RowsAffected != 0 {
			commentList[i].User.IsFollow = true
		}
	}
	return &commentList, nil
}
