package comment

import "github.com/gin-gonic/gin"

type CommentHandler interface {
	CreateCommentHdl(ctx *gin.Context)
	GetCommentsHdl(ctx *gin.Context)
	UpdateCommentHdl(ctx *gin.Context)
	DeleteCommentHdl(ctx *gin.Context)
}
