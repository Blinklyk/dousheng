package response

import "github.com/RaymondCode/simple-demo/model"

type FavoriteListResponse struct {
	Response
	VideoList []model.Video `json:"video_list"`
}
