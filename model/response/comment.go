package response

import (
	"github.com/RaymondCode/simple-demo/pb/rpcComment"
)

type CommentActionResponse struct {
	Response
	Comment *rpcComment.Comment `json:"comment,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []*rpcComment.Comment `json:"comment_list"`
}
