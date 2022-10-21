package user

import (
	"context"
)

type UserRepo interface {
	RegisterUser(ctx context.Context, user *User) (err error)
	GetUserById(ctx context.Context, userId uint64) (result User, err error)
	UpdateUser(ctx context.Context, userId uint64, email string, username string) (result User, err error)
	DeleteUser(ctx context.Context, userId uint64) (err error)
}
