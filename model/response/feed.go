package response

import "github.com/RaymondCode/simple-demo/model"

type FeedResponse struct {
	Response
	VideoList []model.Video `json:"video_list"`
	NextTime  int64         `json:"next_time"`
}
