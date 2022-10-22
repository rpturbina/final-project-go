package comment

import (
	"context"

	"github.com/rpturbina/final-project-go/pkg/domain/message"
)

type CommentUsecase interface {
	CreateCommentSvc(ctx context.Context, comment Comment) (result Comment, errMsg message.ErrorMessage)
	GetCommentsSvc(ctx context.Context) (result []Comment, errMsg message.ErrorMessage)
}
