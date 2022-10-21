package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpturbina/final-project-go/pkg/domain/comment"
	"github.com/rpturbina/final-project-go/pkg/domain/message"
	"github.com/rpturbina/final-project-go/pkg/domain/photo"
)

type PhotoHdlImpl struct {
	photoUsecase photo.PhotoUsecase
}

func (p *PhotoHdlImpl) CreatePhotoHdl(ctx *gin.Context) {
	log.Printf("%T - CreatePhotoHdl is invoked\n", p)
	defer log.Printf("%T - CreatePhotoHdl executed\n", p)

	var inputPhoto photo.Photo

	log.Println("binding body payload from request")
	if err := ctx.ShouldBindJSON(&inputPhoto); err != nil {
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

	log.Println("calling create photo usecase service")
	result, errMsg := p.photoUsecase.CreatePhotoSvc(ctx, inputPhoto)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	if result.Comments == nil {
		result.Comments = []comment.Comment{}
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    01,
		"message": "user has successfully registered",
		"type":    "ACCEPTED",
		"data": gin.H{
			"id":         result.ID,
			"title":      result.Title,
			"caption":    result.Caption,
			"url":        result.Url,
			"created_at": result.CreatedAt,
			"user_id":    result.UserID,
		},
	})
}

func NewPhotoHandler(photoUsecase photo.PhotoUsecase) photo.PhotoHandler {
	return &PhotoHdlImpl{photoUsecase: photoUsecase}
}
