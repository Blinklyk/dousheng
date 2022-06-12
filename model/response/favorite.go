package response

type FavoriteListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}
