package response

import "github.com/RaymondCode/simple-demo/model"

type CommentActionResponse struct {
	Response
	Comment model.Comment `json:"comment,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}
