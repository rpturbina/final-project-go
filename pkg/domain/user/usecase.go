package user

import (
	"context"

	"github.com/rpturbina/final-project-go/pkg/domain/message"
)

type UserUsecase interface {
	RegisterUserSvc(ctx context.Context, user User) (result User, errMsg message.ErrorMessage)
}
