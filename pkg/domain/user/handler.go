package user

import "github.com/gin-gonic/gin"

type UserHandler interface {
	RegisterUserHdl(ctx *gin.Context)
	GetUserByIdHdl(ctx *gin.Context)
}
