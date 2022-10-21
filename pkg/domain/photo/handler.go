package photo

import "github.com/gin-gonic/gin"

type PhotoHandler interface {
	CreatePhotoHdl(ctx *gin.Context)
}
