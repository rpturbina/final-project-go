package v1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rpturbina/final-project-go/pkg/domain/comment"
	"github.com/rpturbina/final-project-go/pkg/domain/message"
)

type CommentHdlImpl struct {
	commentUsecase comment.CommentUsecase
}

func (p *CommentHdlImpl) CreateCommentHdl(ctx *gin.Context) {
	log.Printf("%T - CreateCommentHdl is invoked\n", p)
	defer log.Printf("%T - CreateCommentHdl executed\n", p)

	var inputComment comment.Comment

	log.Println("binding body payload from request")
	if err := ctx.ShouldBindJSON(&inputComment); err != nil {
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

	log.Println("calling create comment usecase service")
	result, errMsg := p.commentUsecase.CreateCommentSvc(ctx, inputComment)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    01,
		"message": "comment has successfully created",
		"type":    "ACCEPTED",
		"data": gin.H{
			"id":         result.ID,
			"message":    result.Message,
			"photo_id":   result.PhotoID,
			"user_id":    result.UserID,
			"created_at": result.CreatedAt,
		},
	})
}

func (p *CommentHdlImpl) GetCommentsHdl(ctx *gin.Context) {
	log.Printf("%T - GetCommentsIdHdl is invoked\n", p)
	defer log.Printf("%T - GetCommentsHdl executed\n", p)

	stringUserId := ctx.Value("user").(string)

	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

	log.Println("calling get comments by user id usecase service")
	result, errMsg := p.commentUsecase.GetCommentsSvc(ctx)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"message": fmt.Sprintf("comments by user id %v is found", userId),
		"type":    "SUCCESS",
		"data":    result,
	})
}

// func (p *CommentHdlImpl) UpdateCommentHdl(ctx *gin.Context) {
// 	log.Printf("%T - UpdateCommentHdl is invoked\n", p)
// 	defer log.Printf("%T - UpdateCommentHdl executed\n", p)

// 	log.Println("check commentId from path parameter")
// 	commentIdParam := ctx.Param("commentId")

// 	commentId, err := strconv.ParseUint(commentIdParam, 0, 64)

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

// 	log.Println("calling get comment by id usecase service")
// 	result, errMsg := p.commentUsecase.GetCommentByIdSvc(ctx, commentId)

// 	if errMsg.Error != nil {
// 		message.ErrorResponseSwitcher(ctx, errMsg)
// 		return
// 	}

// 	stringUserId := ctx.Value("user").(string)
// 	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

// 	log.Println("verify the comment belongs to")
// 	if result.UserID != userId {
// 		message.ErrorResponseSwitcher(ctx, message.ErrorMessage{
// 			Type:  "INVALID_SCOPE",
// 			Error: errors.New("cannot update the comment"),
// 		})
// 		return
// 	}

// 	var updatedComment comment.Comment

// 	log.Println("binding body payload from request")
// 	if err := ctx.ShouldBindJSON(&updatedComment); err != nil {
// 		message.ErrorResponseSwitcher(ctx, message.ErrorMessage{
// 			Type:  "INVALID_PAYLOAD",
// 			Error: errors.New("failed to bind payload"),
// 		})
// 		return
// 	}

// 	ctx.Set("commentId", result.ID)

// 	result, errMsg = p.commentUsecase.UpdateCommentSvc(ctx, updatedComment.Title, updatedComment.Caption, updatedComment.Url)

// 	if errMsg.Error != nil {
// 		message.ErrorResponseSwitcher(ctx, errMsg)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"code":    01,
// 		"message": "comment has been successfully updated",
// 		"type":    "ACCEPTED",
// 		"data":    result,
// 	})
// }

// func (p *CommentHdlImpl) DeleteCommentHdl(ctx *gin.Context) {
// 	log.Printf("%T - DeleteCommentHdl is invoked\n", p)
// 	defer log.Printf("%T - DeleteCommentHdl executed\n", p)

// 	log.Println("check commentId from path parameter")
// 	commentIdParam := ctx.Param("commentId")

// 	commentId, err := strconv.ParseUint(commentIdParam, 0, 64)

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

// 	log.Println("calling get comment by id usecase service")
// 	result, errMsg := p.commentUsecase.GetCommentByIdSvc(ctx, commentId)

// 	if errMsg.Error != nil {
// 		message.ErrorResponseSwitcher(ctx, errMsg)
// 		return
// 	}

// 	stringUserId := ctx.Value("user").(string)
// 	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

// 	log.Println("verify the comment belongs to")
// 	if result.UserID != userId {
// 		message.ErrorResponseSwitcher(ctx, message.ErrorMessage{
// 			Type:  "INVALID_SCOPE",
// 			Error: errors.New("cannot delete the comment"),
// 		})
// 		return
// 	}

// 	log.Println("calling delete comment usecase service")
// 	errMsg = p.commentUsecase.DeleteCommentSvc(ctx, commentId)

// 	if errMsg.Error != nil {
// 		message.ErrorResponseSwitcher(ctx, errMsg)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"code":    01,
// 		"message": "comment has been successfully deleted",
// 		"type":    "ACCEPTED",
// 	})
// }

func NewCommentHandler(commentUsecase comment.CommentUsecase) comment.CommentHandler {
	return &CommentHdlImpl{commentUsecase: commentUsecase}
}
