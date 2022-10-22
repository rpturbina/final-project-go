package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpturbina/final-project-go/pkg/domain/auth"
	"github.com/rpturbina/final-project-go/pkg/domain/message"
	"github.com/rpturbina/final-project-go/pkg/domain/user"
)

type AuthHdlImpl struct {
	authUsecase auth.AuthUsecase
}

func (a *AuthHdlImpl) LoginUserHdl(ctx *gin.Context) {
	log.Printf("%T - LoginUserHdl is invoked\n", a)
	defer log.Printf("%T - LoginUserHdl executed\n", a)

	var user user.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed to bind payload",
			"invalid_arg": gin.H{
				"error_type":    "INVALID_PAYLOAD",
				"error_message": err.Error(),
			},
		})
		return
	}

	accessToken, refreshToken, idToken, errMsg := a.authUsecase.LoginUserSvc(ctx, user.Username, user.Password)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"message": "successfully login",
		"type":    "SUCCESS",
		"data": gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"id_token":      idToken,
		},
	})
}

func (a *AuthHdlImpl) GetRefreshTokenHdl(ctx *gin.Context) {
	log.Printf("%T - GetRefreshTokenHdl is invoked\n", a)
	defer log.Printf("%T - GetRefreshTokenHdl executed\n", a)

	accessToken, refreshToken, idToken, errMsg := a.authUsecase.GetRefreshTokenSvc(ctx)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"message": "token refreshed successfully",
		"type":    "ACCEPTED",
		"data": gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"id_token":      idToken,
		},
	})

}

func NewAuthHandler(authUsecase auth.AuthUsecase) auth.AuthHandler {
	return &AuthHdlImpl{authUsecase: authUsecase}
}
