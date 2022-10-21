package v1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
		"message": "photo has successfully created",
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

func (p *PhotoHdlImpl) GetPhotosByUserIdHdl(ctx *gin.Context) {
	log.Printf("%T - GetPhotosByUserIdHdl is invoked\n", p)
	defer log.Printf("%T - GetPhotosByUserIdHdl executed\n", p)

	a, b := ctx.GetQueryArray("user_id")
	_, _ = a, b

	stringUserId, isExist := ctx.GetQuery("user_id")

	if !isExist {
		stringUserId = ctx.Value("user").(string)
	}

	if isExist && stringUserId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "invalid params",
			"invalid_arg": gin.H{
				"error_type":    "INVALID_PARAMS",
				"error_message": "invalid params",
			},
		})
		return
	}

	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

	log.Println("calling create photo usecase service")
	result, errMsg := p.photoUsecase.GetPhotosByUserIdSvc(ctx, userId)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    00,
		"message": fmt.Sprintf("photos by user id %v is found", userId),
		"type":    "SUCCESS",
		"data":    result,
	})
}

func NewPhotoHandler(photoUsecase photo.PhotoUsecase) photo.PhotoHandler {
	return &PhotoHdlImpl{photoUsecase: photoUsecase}
}
