package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpturbina/final-project-go/pkg/domain/auth"
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
				"error_type":    "INVALID_FORMAT",
				"error_message": err.Error(),
			},
		})
		return
	}

	accessToken, refreshToken, idToken, errMsg := a.authUsecase.LoginUserSvc(ctx, user.Username, user.Password)

	if errMsg.Error != nil {
		switch errMsg.Type {
		case "INTERNAL_CONNECTION_PROBLEM":
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code": 99,
				"type": "INTERNAL_SERVER_ERROR",
				"invalid_arg": gin.H{
					"error_type":    errMsg.Type,
					"error_message": errMsg.Error.Error(),
				},
			})
		case "WRONG_PASSWORD":
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code": 97,
				"type": "UNAUTHENTICATED",
				"invalid_arg": gin.H{
					"error_type":    errMsg.Type,
					"error_message": errMsg.Error.Error(),
				},
			})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    01,
		"message": "successfully login",
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