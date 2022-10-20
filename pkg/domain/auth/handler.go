package auth

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	LoginUserHdl(ctx *gin.Context)
}
