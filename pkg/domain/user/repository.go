package user

import "context"

type UserRepo interface {
	RegisterUser(ctx context.Context, user *User) (err error)
	GetUserById(ctx context.Context, userId uint64) (result User, err error)
}
