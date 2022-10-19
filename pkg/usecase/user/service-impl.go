package user

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/rpturbina/final-project-go/pkg/domain/message"
	"github.com/rpturbina/final-project-go/pkg/domain/user"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func (u *UserUsecaseImpl) RegisterUserSvc(ctx context.Context, user user.User) (result user.User, errMsg message.ErrorMessage) {
	log.Printf("%T - RegisterUserSvc is invoked\n", u)
	defer log.Printf("%T - RegisterUserSvc executed\n", u)

	// TODO: creating error when username and email is already exist

	log.Println("create user to database")
	err := u.userRepo.RegisterUser(ctx, &user)
	if err != nil {
		log.Printf("error when creating user: %v\n", err.Error())

		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "idx_users_user_name"`) {
			err = errors.New("username has already been registered")
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "USER_REGISTERED",
			}
			return user, errMsg
		}

		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "idx_users_email"`) {
			err = errors.New("email has already been registered")
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "USER_REGISTERED",
			}

			return user, errMsg
		}
	}

	return user, errMsg
}

func NewUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo}
}
