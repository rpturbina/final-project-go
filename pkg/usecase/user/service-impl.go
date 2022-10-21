package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
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

	// user input validation
	if isValid, err := govalidator.ValidateStruct(user); !isValid {
		switch err.Error() {
		case "username is required":
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "USERNAME_IS_EMPTY",
			}
			return result, errMsg
		case "email is required":
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "EMAIL_IS_EMPTY",
			}
			return result, errMsg
		case "invalid email format":
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "WRONG_EMAIL_FORMAT",
			}
			return result, errMsg
		case "password is required":
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "PASSWORD_IS_EMPTY",
			}
			return result, errMsg
		case "password has to have a minimum length of 6 characters":
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "INVALID_PASSWORD_FORMAT",
			}
			return result, errMsg

		default:
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "INVALID_PAYLOAD",
			}
			return result, errMsg
		}
	}

	log.Println("calling register user repo")
	err := u.userRepo.RegisterUser(ctx, &user)
	if err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "idx_users_username"`) {
			err = errors.New("username has already been registered")
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "USERNAME_REGISTERED",
			}
			return result, errMsg
		}

		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "idx_users_email"`) {
			err = errors.New("email has already been registered")
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "EMAIL_REGISTERED",
			}

			return result, errMsg
		}

	}

	if err != nil {
		log.Printf("error when fetching data from database: %s\n", err.Error())
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "INTERNAL_CONNECTION_PROBLEM",
		}
		return result, errMsg
	}

	return user, errMsg
}

func (u *UserUsecaseImpl) GetUserByIdSvc(ctx context.Context, userId uint64) (result user.User, errMsg message.ErrorMessage) {
	log.Printf("%T - GetUserByIdSvc is invoked\n", u)
	defer log.Printf("%T - GetUserByIdSvc executed\n", u)

	log.Println("calling get user by id repo")
	result, err := u.userRepo.GetUserById(ctx, userId)

	if err != nil {
		log.Printf("error when fetching data from database: %s\n", err.Error())
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "INTERNAL_CONNECTION_PROBLEM",
		}
		return result, errMsg
	}

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

func (u *UserUsecaseImpl) UpdateUserSvc(ctx context.Context, userId uint64, email string, username string) (idToken string, errMsg message.ErrorMessage) {
	log.Printf("%T - UpdateUserSvc is invoked\n", u)
	defer log.Printf("%T - UpdateUserSvc executed\n", u)

	if email == "" {
		errMsg := message.ErrorMessage{
			Error: errors.New("email is required"),
			Type:  "EMAIL_IS_EMPTY",
		}
		return idToken, errMsg
	}

	// email validation
	if !govalidator.IsEmail(email) {
		errMsg := message.ErrorMessage{
			Error: errors.New("invalid email format"),
			Type:  "WRONG_EMAIL_FORMAT",
		}
		return idToken, errMsg
	}

	// username validation
	if username == "" {
		errMsg := message.ErrorMessage{
			Error: errors.New("username is required"),
			Type:  "USERNAME_IS_EMPTY",
		}
		return idToken, errMsg
	}

	result, err := u.userRepo.UpdateUser(ctx, userId, email, username)

	if err != nil {
		log.Printf("error when fetching data from database: %s\n", err.Error())
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "INTERNAL_CONNECTION_PROBLEM",
		}
		return idToken, errMsg
	}

	claimId := claim.IDToken{
		JWTID:    uuid.New(),
		Username: result.Username,
		Email:    result.Email,
		DOB:      time.Time(result.DOB),
	}
	idToken, _ = crypto.CreateJWT(ctx, claimId)

	return idToken, errMsg
}

func (u *UserUsecaseImpl) DeleteUserSvc(ctx context.Context, userId uint64) (errMsg message.ErrorMessage) {
	log.Printf("%T - UpdateUserSvc is invoked\n", u)
	defer log.Printf("%T - UpdateUserSvc executed\n", u)

	log.Println("calling delete user repo")
	err := u.userRepo.DeleteUser(ctx, userId)

	if err != nil {
		log.Printf("error when fetching data from database: %s\n", err.Error())
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "INTERNAL_CONNECTION_PROBLEM",
		}
		return errMsg
	}

	return errMsg
}

func NewUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo}
}
