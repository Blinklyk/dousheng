package response

type PublishListResponse struct {
	Response
	VideoList []VideoDTO `json:"video_list"`
}
