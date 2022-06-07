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

type FavoriteService struct{}

// FavoriteAction two favorite action type: add favorite or cancel favorite
func (fs *FavoriteService) FavoriteAction(u *model.User, r *request.FavoriteRequest) error {
	userID := u.ID
	videoID := r.VideoID
	actionType := r.ActionType

	// converse to int64 format
	// TODO verify userID nad videoID
	videoIDNum, _ := strconv.ParseInt(videoID, 10, 64)

	// action_type determines operationCount

	favoriteInfo := &model.Favorite{
		UserID:  userID,
		VideoID: videoIDNum,
	}

	if actionType == "1" {

		// check if add already
		res := global.DY_DB.Where("user_id = ? AND video_id = ?", userID, videoID).First(&model.Favorite{})
		if res.RowsAffected != 0 {
			return errors.New("already add favorite this video")
		}

		// add favorite transaction:
		// 1. create data in user_favorite_video table
		// 2. update favorite_count in videos table
		AddFavorite := func(x *gorm.DB) error {
			tx := global.DY_DB.Begin()

			if err := tx.Create(&favoriteInfo).Error; err != nil {
				tx.Rollback()
				log.Println("error when insert u_f_v :", err)
				return errors.New("error: when insert favorite")
			}

			// update the favorite_count column (lock)
			if res := tx.Model(&model.Video{}).Where("id = ?", videoID).
				Clauses(clause.Locking{Strength: "UPDATE"}).
				UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)); res.RowsAffected == 0 {
				tx.Rollback()
				return errors.New("res RowsAffected in videos is 0")
			}

			err := tx.Commit().Error
			return err
		}

		err := AddFavorite(global.DY_DB)
		if err != nil {
			return errors.New("error when adding favorite: " + err.Error())
		}
	}

	if actionType == "2" {
		// check if delete already
		res := global.DY_DB.Where("user_id = ? AND video_id = ?", userID, videoID).First(&model.Favorite{})
		if res.RowsAffected == 0 {
			return errors.New("err: No add favorite this video before")
		}

		// cancel favorite transaction:
		// 1. delete data in user_favorite_video table  (soft delete)
		// 2. update favorite_count in videos table
		CancelFavorite := func(x *gorm.DB) error {
			tx := global.DY_DB.Begin()

			if err := tx.Delete(&model.Favorite{}, "user_id = ? AND video_id = ?", userID, videoIDNum).Error; err != nil {
				log.Println("error when delete u_f_v :", err)
				tx.Rollback()
				return err
			}

			if res := tx.Model(&model.Video{}).Where("id = ?", videoID).
				Clauses(clause.Locking{Strength: "UPDATE"}).
				UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", -1)); res.RowsAffected == 0 {
				tx.Rollback()
				return errors.New("RowsAffected in video table is 0")
			}

			err := tx.Commit().Error
			return err
		}

		err := CancelFavorite(global.DY_DB)
		if err != nil {
			return errors.New("error when canceling favorite: " + err.Error())
		}
	}
	return nil
}

// FavoriteList get the videos that user favorites
func (fs *FavoriteService) FavoriteList(r *request.FavoriteListRequest) (favoriteVideoList *[]model.Video, err error) {
	userID := r.UserID
	// find favorite videos from db
	var videosID []int64
	// get video_id from conn table first
	// use distinct instead of select, and check if the deleted_at column is null
	if err := global.DY_DB.Table("dy_favorite").Distinct("video_id").Where("user_id = ? AND deleted_at is null", userID).Find(&videosID).Error; err != nil {
		return nil, err
	}

	// get video details from video table by selecting video_id
	if err := global.DY_DB.Model(&model.Video{}).Where("ID in ?", videosID).Preload("User").Find(&favoriteVideoList).Error; err != nil {
		return nil, errors.New("get videos from derived videosID" + err.Error())
	}

	// add is_favorite and is_follow info to the video list
	userIDNum, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, errors.New("error: strconv userId to 64")
	}
	returnFavoriteVideoList, err := VideoListAppendInfo(*favoriteVideoList, userIDNum)
	if err != nil {
		return nil, err
	}
	return returnFavoriteVideoList, nil
}
