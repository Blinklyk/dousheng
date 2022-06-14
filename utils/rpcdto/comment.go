package rpcdto

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/pb/rpcComment"
)

// ToCommentRpcDTO  model comment to rpc comment
func ToCommentRpcDTO(comment *model.Comment) *rpcComment.Comment {
	commentRpcDTO := &rpcComment.Comment{
		Id:         comment.ID,
		User:       ToUserRpcDTO(&comment.User),
		Content:    comment.Content,
		CreateDate: comment.CreateData,
	}
	return commentRpcDTO
}

// ToCommentListRpcDTO model comment list to rpc comment list
func ToCommentListRpcDTO(comments []model.Comment) []*rpcComment.Comment {

	commentsReturn := make([]*rpcComment.Comment, len(comments), len(comments))
	// traverse the comment list and add c to commentReturn list
	for i := 0; i < len(comments); i++ {
		c := &rpcComment.Comment{
			Id:         comments[i].ID,
			User:       ToUserRpcDTO(&comments[i].User),
			Content:    comments[i].Content,
			CreateDate: comments[i].CreateData,
		}

		commentsReturn[i] = c
	}
	return commentsReturn
}
