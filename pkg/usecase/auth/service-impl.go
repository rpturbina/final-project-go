package auth

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/rpturbina/final-project-go/helpers"
	"github.com/rpturbina/final-project-go/pkg/domain/auth"
	"github.com/rpturbina/final-project-go/pkg/domain/claim"
	"github.com/rpturbina/final-project-go/pkg/domain/message"
	"github.com/rpturbina/final-project-go/pkg/usecase/crypto"
)

type AuthUsecaseImpl struct {
	authRepo auth.AuthRepo
}

func (a *AuthUsecaseImpl) LoginUserSvc(ctx context.Context, username string, password string) (accessToken string, refreshToken string, idToken string, errMsg message.ErrorMessage) {
	log.Printf("%T - LoginUserSvc is invoked\n", a)
	defer log.Printf("%T - LoginUserSvc executed\n", a)

	result, err := a.authRepo.LoginUser(ctx, username)

	if err != nil {
		log.Printf("error when fetching data from database: %s\n", err.Error())
		return accessToken, refreshToken, idToken, errMsg
	}

	comparePass := helpers.ComparePass(
		[]byte(result.Password), []byte(password),
	)

	if !comparePass {
		err := errors.New("invalid email or password")
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "WRONG_PASSWORD",
		}
		return accessToken, refreshToken, idToken, errMsg
	}

	timeNow := time.Now()
	claimAccess := claim.JWTToken{
		JWTID:          uuid.New(),
		Subject:        result.Username,
		Issuer:         "mygram.com",
		Audience:       "user.mygram.com",
		Scope:          "create update delete read",
		Type:           "ACCESS_TOKEN",
		IssuedAt:       timeNow.Unix(),
		NotValidBefore: timeNow.Unix(),
		ExpiredAt:      timeNow.Add(24 * time.Hour).Unix(),
	}

	accessToken, _ = crypto.CreateJWT(ctx, claimAccess)

	claimRefresh := claim.JWTToken{
		JWTID:          uuid.New(),
		Subject:        result.Username,
		Issuer:         "mygram.com",
		Audience:       "user.mygram.com",
		Scope:          "create update delete read",
		Type:           "REFRESH_TOKEN",
		IssuedAt:       timeNow.Unix(),
		NotValidBefore: timeNow.Unix(),
		ExpiredAt:      timeNow.Add(1000 * time.Hour).Unix(),
	}
	refreshToken, _ = crypto.CreateJWT(ctx, claimRefresh)

	claimId := claim.IDToken{
		JWTID:    uuid.New(),
		Username: result.Username,
		Email:    result.Email,
		DOB:      time.Time(result.DOB),
	}
	idToken, _ = crypto.CreateJWT(ctx, claimId)

	return accessToken, refreshToken, idToken, errMsg
}

func NewAuthUsecase(authRepo auth.AuthRepo) auth.AuthUsecase {
	return &AuthUsecaseImpl{authRepo: authRepo}
}
