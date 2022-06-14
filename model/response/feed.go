package response

import (
	"github.com/RaymondCode/simple-demo/model/dto"
)

type FeedResponse struct {
	Response
	VideoList *[]dto.VideoDTO `json:"video_list"`
	NextTime  int64           `json:"next_time"`
}
