package photo

import (
	"context"

	"github.com/rpturbina/final-project-go/pkg/domain/message"
)

type PhotoUsecase interface {
	CreatePhotoSvc(ctx context.Context, photo Photo) (result Photo, errMsg message.ErrorMessage)
	GetPhotosByUserIdSvc(ctx context.Context, userId uint64) (result []Photo, errMsg message.ErrorMessage)
}