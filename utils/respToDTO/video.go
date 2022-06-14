package respToDTO

import (
	"github.com/RaymondCode/simple-demo/model/dto"
	"github.com/RaymondCode/simple-demo/pb/rpcVideo"
)

func GetVideoDTo(video *rpcVideo.Video) *dto.VideoDTO {
	var videoInfo dto.VideoDTO
	videoInfo.ID = video.Id
	videoInfo.UserID = video.User.Id
	videoInfo.User = *GetUserDTo(video.User)
	videoInfo.PlayUrl = video.PlayUrl
	videoInfo.CoverUrl = video.CoverUrl
	videoInfo.FavoriteCount = video.FavoriteCount
	videoInfo.CommentCount = video.CommentCount
	videoInfo.IsFavorite = video.IsFavorite
	videoInfo.Title = video.Title
	return &videoInfo
}

func GetVideoListDTo(videos []*rpcVideo.Video) *[]dto.VideoDTO {
	videoInfo := make([]dto.VideoDTO, len(videos))
	for i := 0; i < len(videos); i++ {
		videoInfo[i] = *GetVideoDTo(videos[i])
	}
	return &videoInfo
}
