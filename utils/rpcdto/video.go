package rpcdto

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/pb/rpcVideo"
)

// ToVideoListRpcDTO model video list to rpc video list and nested rpc user
func ToVideoListRpcDTO(videos []model.Video) []*rpcVideo.Video {

	videosReturn := make([]*rpcVideo.Video, len(videos), len(videos))
	// traverse the video list and add v to videosReturn list
	for i := 0; i < len(videos); i++ {
		v := &rpcVideo.Video{
			Id:            videos[i].ID,
			User:          ToUserRpcDTO(&videos[i].User),
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount:  videos[i].CommentCount,
			IsFavorite:    videos[i].IsFavorite,
			Title:         videos[i].Title,
		}

		videosReturn[i] = v

	}
	return videosReturn
}
