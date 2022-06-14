package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/utils"
	"go.uber.org/zap"
	"time"
)

type PublishService struct{}

func (ps *PublishService) PublishAction(u *model.User, r *request.PublishRequest, filePath string) error {
	title := r.Title

	// upload the file to oss and get the url from oss
	videoUrl, coverUrl, err := utils.UploadFile(filePath)
	if err != nil {
		global.App.DY_LOG.Error("upload video error!", zap.Error(err))
		return err
	}

	publishVideo := &model.Video{
		UserID:        u.ID,
		PlayUrl:       videoUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		PublishTime:   time.Now(),
		Title:         title,
		IsFavorite:    false,
	}

	if result := global.App.DY_DB.Model(&model.Video{}).Create(&publishVideo); result.RowsAffected == 0 {
		return errors.New("publish error")
	}
	return nil
}

// PublishList return the publishing video list
func (ps *PublishService) PublishList(r *request.PublishListRequest) (publishVideos []model.Video, err error) {
	if err := global.App.DY_DB.Where("user_id = ?", r.UserID).Preload("User").Order("ID desc").Find(&publishVideos).Error; err != nil {
		return nil, err
	}
	// add is_favorite and is_follow value
	//userIDNum, err := strconv.ParseInt(r.UserID, 10, 64)
	if err != nil {
		return nil, errors.New("error: conv userID to int64 ")
	}
	VideoListAppendInfo(publishVideos, r.UserID)
	return
}
