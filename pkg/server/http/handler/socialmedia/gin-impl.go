package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpturbina/final-project-go/pkg/domain/message"
	"github.com/rpturbina/final-project-go/pkg/domain/socialmedia"
)

type SocialMediaHdlImpl struct {
	socialMediaUsecase socialmedia.SocialMediaUsecase
}

func (c *SocialMediaHdlImpl) CreateSocialMediaHdl(ctx *gin.Context) {
	log.Printf("%T - CreateSocialMediaHdl is invoked\n", c)
	defer log.Printf("%T - CreateSocialMediaHdl executed\n", c)

	var inputSocialMedia socialmedia.SocialMedia

	log.Println("binding body payload from request")
	if err := ctx.ShouldBindJSON(&inputSocialMedia); err != nil {
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

	log.Println("calling create socialMedia usecase service")
	result, errMsg := c.socialMediaUsecase.CreateSocialMediaSvc(ctx, inputSocialMedia)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    01,
		"message": "social media has successfully created",
		"type":    "ACCEPTED",
		"data": gin.H{
			"id":         result.ID,
			"name":       result.Name,
			"url":        result.URL,
			"created_at": result.CreatedAt,
			"user_id":    result.UserID,
		},
	})
}

// func (c *SocialMediaHdlImpl) GetSocialMediasHdl(ctx *gin.Context) {
// 	log.Printf("%T - GetSocialMediasIdHdl is invoked\n", c)
// 	defer log.Printf("%T - GetSocialMediasHdl executed\n", c)

// 	stringUserId := ctx.Value("user").(string)

// 	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

// 	log.Println("calling get socialMedias by user id usecase service")
// 	result, errMsg := c.socialMediaUsecase.GetSocialMediasSvc(ctx)

// 	if errMsg.Error != nil {
// 		message.ErrorResponseSwitcher(ctx, errMsg)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"code":    00,
// 		"message": fmt.Sprintf("socialMedias by user id %v is found", userId),
// 		"type":    "SUCCESS",
// 		"data":    result,
// 	})
// }

// func (c *SocialMediaHdlImpl) UpdateSocialMediaHdl(ctx *gin.Context) {
// 	log.Printf("%T - UpdateSocialMediaHdl is invoked\n", c)
// 	defer log.Printf("%T - UpdateSocialMediaHdl executed\n", c)

// 	log.Println("check socialMediaId from path parameter")
// 	socialMediaIdParam := ctx.Param("socialMediaId")

// 	socialMediaId, err := strconv.ParseUint(socialMediaIdParam, 0, 64)

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"code":    96,
// 			"type":    "BAD_REQUEST",
// 			"message": "invalid params",
// 			"invalid_arg": gin.H{
// 				"error_type":    "INVALID_PARAMS",
// 				"error_message": "invalid params",
// 			},
// 		})
// 		return
// 	}

// 	log.Println("calling get socialMedia by id usecase service")
// 	result, errMsg := c.socialMediaUsecase.GetSocialMediaByIdSvc(ctx, socialMediaId)

// 	if errMsg.Error != nil {
// 		message.ErrorResponseSwitcher(ctx, errMsg)
// 		return
// 	}

// 	stringUserId := ctx.Value("user").(string)
// 	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

// 	log.Println("verify the socialMedia belongs to")
// 	if result.UserID != userId {
// 		message.ErrorResponseSwitcher(ctx, message.ErrorMessage{
// 			Type:  "INVALID_SCOPE",
// 			Error: errors.New("cannot update the socialMedia"),
// 		})
// 		return
// 	}

// 	var updatedSocialMedia socialMedia.SocialMedia

// 	log.Println("binding body payload from request")
// 	if err := ctx.ShouldBindJSON(&updatedSocialMedia); err != nil {
// 		message.ErrorResponseSwitcher(ctx, message.ErrorMessage{
// 			Type:  "INVALID_PAYLOAD",
// 			Error: errors.New("failed to bind payload"),
// 		})
// 		return
// 	}

// 	ctx.Set("socialMediaId", socialMediaId)

// 	updateResult, errMsg := c.socialMediaUsecase.UpdateSocialMediaSvc(ctx, updatedSocialMedia.Message)

// 	if errMsg.Error != nil {
// 		message.ErrorResponseSwitcher(ctx, errMsg)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"code":    01,
// 		"message": "socialMedia has been successfully updated",
// 		"type":    "ACCEPTED",
// 		"data":    updateResult,
// 	})
// }

// func (c *SocialMediaHdlImpl) DeleteSocialMediaHdl(ctx *gin.Context) {
// 	log.Printf("%T - DeleteSocialMediaHdl is invoked\n", c)
// 	defer log.Printf("%T - DeleteSocialMediaHdl executed\n", c)

// 	log.Println("check socialMediaId from path parameter")
// 	socialMediaIdParam := ctx.Param("socialMediaId")

// 	socialMediaId, err := strconv.ParseUint(socialMediaIdParam, 0, 64)

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"code":    96,
// 			"type":    "BAD_REQUEST",
// 			"message": "invalid params",
// 			"invalid_arg": gin.H{
// 				"error_type":    "INVALID_PARAMS",
// 				"error_message": "invalid params",
// 			},
// 		})
// 		return
// 	}

// 	log.Println("calling get socialMedia by id usecase service")
// 	result, errMsg := c.socialMediaUsecase.GetSocialMediaByIdSvc(ctx, socialMediaId)

// 	if errMsg.Error != nil {
// 		message.ErrorResponseSwitcher(ctx, errMsg)
// 		return
// 	}

// 	stringUserId := ctx.Value("user").(string)
// 	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

// 	log.Println("verify the socialMedia belongs to")
// 	if result.UserID != userId {
// 		message.ErrorResponseSwitcher(ctx, message.ErrorMessage{
// 			Type:  "INVALID_SCOPE",
// 			Error: errors.New("cannot delete the socialMedia"),
// 		})
// 		return
// 	}

// 	log.Println("calling delete socialMedia usecase service")
// 	errMsg = c.socialMediaUsecase.DeleteSocialMediaSvc(ctx, socialMediaId)

// 	if errMsg.Error != nil {
// 		message.ErrorResponseSwitcher(ctx, errMsg)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"code":    01,
// 		"message": "socialMedia has been successfully deleted",
// 		"type":    "ACCEPTED",
// 	})
// }

func NewSocialMediaHandler(socialMediaUsecase socialmedia.SocialMediaUsecase) socialmedia.SocialMediaHandler {
	return &SocialMediaHdlImpl{socialMediaUsecase: socialMediaUsecase}
}
