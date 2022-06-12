package response

type FeedResponse struct {
	Response
	VideoList []VideoDTO `json:"video_list"`
	NextTime  int64      `json:"next_time"`
}
