package request

type FeedRequest struct {
	LatestTime string `json:"latest_time,omitempty"`
	Token      string `json:"token" form:"token"`
}
