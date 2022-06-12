package service

import (
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
)

// VideoListAppendInfo add is_follow and is_favorite value to a videoList get from db
func VideoListAppendInfo(vs []model.Video, userID int64) (*[]model.Video, error) {
	for i := 0; i < len(vs); i++ {
		// determine is_favorite value
		var tmp model.Favorite
		if res := global.App.DY_DB.Model(&model.Favorite{}).Where("user_id = ? AND video_id = ?", userID, vs[i].ID).First(&tmp); res.RowsAffected != 0 {
			vs[i].IsFavorite = true
		}
		// determine is_follow value
		var temp model.Follow
		if res := global.App.DY_DB.Model(&model.Follow{}).Where("user_id = ? AND follow_id = ?", userID, vs[i].UserID).First(&temp); res.RowsAffected != 0 {
			vs[i].User.IsFollow = true
		}
	}
	return &vs, nil
}
