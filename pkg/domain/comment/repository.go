package comment

import (
	"context"
)

type CommentRepo interface {
	CreateComment(ctx context.Context, inputComment *Comment) (result Comment, err error)
	GetComments(ctx context.Context, userId uint64) (result []Comment, err error)
	GetCommentById(ctx context.Context, commentId uint64) (result Comment, err error)
	UpdateComment(ctx context.Context, inputMessage string) (result Comment, err error)
	DeleteComment(ctx context.Context, commentId uint64) (err error)
}
