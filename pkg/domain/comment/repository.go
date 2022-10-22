package comment

import "context"

type CommentRepo interface {
	CreateComment(ctx context.Context, inputComment *Comment) (result Comment, err error)
	GetComments(ctx context.Context, userId uint64) (result []Comment, err error)
}
