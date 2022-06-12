package response

type PublishListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}
