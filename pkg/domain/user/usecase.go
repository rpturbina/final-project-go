package user

import (
	"context"

	"github.com/rpturbina/final-project-go/pkg/domain/message"
)

type UserUsecase interface {
	RegisterUserSvc(ctx context.Context, user User) (result User, errMsg message.ErrorMessage)
	GetUserByIdSvc(ctx context.Context, userId uint64) (result User, errMsg message.ErrorMessage)
}
