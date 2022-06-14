package response

import (
	"github.com/RaymondCode/simple-demo/model/dto"
)

type PublishListResponse struct {
	Response
	VideoList []dto.VideoDTO `json:"video_list"`
}
