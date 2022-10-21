package comment

import "context"

type CommentRepo interface {
	CreateComment(ctx context.Context, inputComment *Comment) (result Comment, err error)
}
