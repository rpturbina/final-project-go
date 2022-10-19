package user

import (
	"context"
	"log"

	"github.com/rpturbina/final-project-go/config/postgres"
	"github.com/rpturbina/final-project-go/pkg/domain/user"
)

type UserRepoImpl struct {
	pgCln postgres.PostgresClient
}

func (u *UserRepoImpl) RegisterUser(ctx context.Context, inputUser *user.User) (err error) {
	log.Printf("%T - RegisterUser is invoked\n", u)
	defer log.Printf("%T - RegisterUser executed\n", u)

	db := u.pgCln.GetClient()

	err = db.Model(&user.User{}).Create(&inputUser).Error

	if err != nil {
		log.Printf("error when creating user %v\n", inputUser.Email)
	}

	return err
}

func (u *UserRepoImpl) GetUserById(ctx context.Context, userId uint64) (result user.User, err error) {
	log.Printf("%T - GetUserById is invoked\n", u)
	defer log.Printf("%T - GetUserById executed\n", u)

	db := u.pgCln.GetClient()

	err = db.Model(&user.User{}).Where("id = ?", userId).Find(&result).Error

	if err != nil {
		log.Printf("error when getting user by id %v\n", userId)
	}

	return result, err
}

func NewUserRepo(pgCln postgres.PostgresClient) user.UserRepo {
	return &UserRepoImpl{pgCln: pgCln}
}
