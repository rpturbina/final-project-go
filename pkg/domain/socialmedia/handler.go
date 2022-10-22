package socialmedia

import "github.com/gin-gonic/gin"

type SocialMediaHandler interface {
	CreateSocialMediaHdl(ctx *gin.Context)
	GetSocialMediasHdl(ctx *gin.Context)
	UpdateSocialMediaHdl(ctx *gin.Context)
	DeleteSocialMediaHdl(ctx *gin.Context)
}
