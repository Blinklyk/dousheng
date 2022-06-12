package response

type FavoriteListResponse struct {
	Response
	VideoList []VideoDTO `json:"video_list"`
}
