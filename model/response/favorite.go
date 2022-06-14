package response

import (
	"github.com/RaymondCode/simple-demo/model/dto"
)

type FavoriteListResponse struct {
	Response
	VideoList []dto.VideoDTO `json:"video_list"`
}
