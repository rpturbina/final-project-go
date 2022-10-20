package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rpturbina/final-project-go/pkg/domain/claim"
	"github.com/rpturbina/final-project-go/pkg/domain/message"
	"github.com/rpturbina/final-project-go/pkg/domain/user"
	"github.com/rpturbina/final-project-go/pkg/usecase/crypto"
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

func (u *UserUsecaseImpl) GetUserByIdSvc(ctx context.Context, userId uint64) (result user.User, errMsg message.ErrorMessage) {
	log.Printf("%T - GetUserByIdSvc is invoked\n", u)
	defer log.Printf("%T - GetUserByIdSvc executed\n", u)

	log.Println("getting user from user repository")
	result, err := u.userRepo.GetUserById(ctx, userId)

	if err != nil {
		log.Printf("error when fetching data from database: %s\n", err.Error())
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "INTERNAL_CONNECTION_PROBLEM",
		}
		return result, errMsg
	}

	log.Println("checking user id")
	if result.ID <= 0 {
		log.Printf("user with id %v is not found", userId)

		err = fmt.Errorf("user with id %v is not found", userId)
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "USER_NOT_FOUND",
		}
		return result, errMsg
	}

	return result, errMsg
}

func (u *UserUsecaseImpl) UpdateUserByIdSvc(ctx context.Context, userId uint64, email string, username string) (idToken string, errMsg message.ErrorMessage) {
	log.Printf("%T - UpdateUserByIdSvc is invoked\n", u)
	defer log.Printf("%T - UpdateUserByIdSvc executed\n", u)

	result, err := u.GetUserByIdSvc(ctx, userId)

	if err.Error != nil {
		log.Printf("error when fetching data from database: %s\n", err.Error.Error())
		errMsg = message.ErrorMessage{
			Error: err.Error,
			Type:  "INTERNAL_CONNECTION_PROBLEM",
		}
		return idToken, errMsg
	}

	log.Println("checking user id")
	if result.ID <= 0 {
		log.Printf("user with id %v is not found", userId)

		err.Error = fmt.Errorf("user with id %v is not found", userId)
		errMsg = message.ErrorMessage{
			Error: err.Error,
			Type:  "USER_NOT_FOUND",
		}
		return idToken, errMsg
	}

	result, err.Error = u.userRepo.UpdateUserById(ctx, userId, email, username)

	if err.Error != nil {
		log.Printf("error when fetching data from database: %s\n", err.Error.Error())
		errMsg = message.ErrorMessage{
			Error: err.Error,
			Type:  "INTERNAL_CONNECTION_PROBLEM",
		}
		return idToken, errMsg
	}

	result, err.Error = u.userRepo.GetUserById(ctx, userId)

	claimId := claim.IDToken{
		JWTID:    uuid.New(),
		Username: result.Username,
		Email:    result.Email,
		DOB:      time.Time(result.DOB),
	}
	idToken, _ = crypto.CreateJWT(ctx, claimId)

	return idToken, errMsg
}

func NewUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo}
}
