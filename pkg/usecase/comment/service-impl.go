package comment

import (
	"context"
	"log"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/rpturbina/final-project-go/pkg/domain/comment"
	"github.com/rpturbina/final-project-go/pkg/domain/message"
	"github.com/rpturbina/final-project-go/pkg/domain/photo"
)

type CommentUsecaseImpl struct {
	commentRepo  comment.CommentRepo
	photoUsecase photo.PhotoUsecase
}

func (c *CommentUsecaseImpl) CreateCommentSvc(ctx context.Context, comment comment.Comment) (result comment.Comment, errMsg message.ErrorMessage) {
	log.Printf("%T - CreateCommentSvc is invoked\n", c)
	defer log.Printf("%T - CreateCommentSvc executed\n", c)

	if isValid, err := govalidator.ValidateStruct(comment); !isValid {
		switch err.Error() {
		case "comment message is required":
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "COMMENT_MESSAGE_IS_EMPTY",
			}
			return result, errMsg
		case "photo id is required":
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "PHOTO_ID_IS_EMPTY",
			}
			return result, errMsg
		default:
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "INVALID_FORMAT",
			}
			return result, errMsg
		}
	}

	log.Println("check photo is exist or not")
	_, errMsg = c.photoUsecase.GetPhotoByIdSvc(ctx, comment.PhotoID)

	if errMsg.Error != nil {
		return result, errMsg
	}

	stringUserId := ctx.Value("user").(string)

	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

	comment.UserID = userId

	log.Println("calling create comment repo")
	result, err := c.commentRepo.CreateComment(ctx, &comment)

	if err != nil {
		log.Printf("error when fetching data from database: %s\n", err.Error())
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "INTERNAL_CONNECTION_PROBLEM",
		}
		return result, errMsg
	}
	return result, errMsg
}

// func (c *CommentUsecaseImpl) GetCommentsByUserIdSvc(ctx context.Context, userId uint64) (result []comment.Comment, errMsg message.ErrorMessage) {
// 	log.Printf("%T - GetCommentsByUserIdSvc is invoked\n", c)
// 	defer log.Printf("%T - GetCommentsByUserIdSvc executed\n", c)

// 	log.Println("calling get user by id repo")
// 	checkUserId, errMsg := c.userUsecase.GetUserByIdSvc(ctx, userId)

// 	if checkUserId.ID <= 0 {
// 		return result, errMsg
// 	}

// 	log.Println("calling get comment by userid repo")
// 	result, err := c.commentRepo.GetCommentsByUserId(ctx, userId)

// 	if err != nil {
// 		log.Printf("error when fetching data from database: %s\n", err.Error())
// 		errMsg = message.ErrorMessage{
// 			Error: err,
// 			Type:  "INTERNAL_CONNECTION_PROBLEM",
// 		}
// 		return result, errMsg
// 	}

// 	return result, errMsg
// }

// func (c *CommentUsecaseImpl) GetCommentByIdSvc(ctx context.Context, commentId uint64) (result comment.Comment, errMsg message.ErrorMessage) {
// 	log.Printf("%T - GetCommentByIdSvc is invoked\n", c)
// 	defer log.Printf("%T - GetCommentByIdSvc executed\n", c)

// 	log.Println("calling get comment by id repo")
// 	result, err := c.commentRepo.GetCommentById(ctx, commentId)

// 	if result.ID <= 0 {
// 		log.Printf("comment with id %v not found", commentId)

// 		err = fmt.Errorf("comment with id %v not found", commentId)
// 		errMsg = message.ErrorMessage{
// 			Error: err,
// 			Type:  "PHOTO_NOT_FOUND",
// 		}
// 		return result, errMsg
// 	}

// 	if err != nil {
// 		log.Printf("error when fetching data from database: %s\n", err.Error())
// 		errMsg = message.ErrorMessage{
// 			Error: err,
// 			Type:  "INTERNAL_CONNECTION_PROBLEM",
// 		}
// 		return result, errMsg
// 	}

// 	return result, errMsg
// }

// func (c *CommentUsecaseImpl) UpdateCommentSvc(ctx context.Context, title string, caption string, url string) (result comment.Comment, errMsg message.ErrorMessage) {
// 	log.Printf("%T - UpdateCommentSvc is invoked\n", c)
// 	defer log.Printf("%T - UpdateCommentSvc executed\n", c)

// 	if title == "" {
// 		errMsg := message.ErrorMessage{
// 			Error: errors.New("comment title is required"),
// 			Type:  "PHOTO_TITLE_IS_EMPTY",
// 		}
// 		return result, errMsg
// 	}
// 	if url == "" {
// 		errMsg := message.ErrorMessage{
// 			Error: errors.New("comment url is required"),
// 			Type:  "PHOTO_URL_IS_EMPTY",
// 		}
// 		return result, errMsg
// 	}

// 	if !govalidator.IsURL(url) {
// 		errMsg := message.ErrorMessage{
// 			Error: errors.New("invalid url format"),
// 			Type:  "INVALID_PHOTO_URL_FORMAT",
// 		}
// 		return result, errMsg
// 	}

// 	result, err := c.commentRepo.UpdateComment(ctx, title, caption, url)

// 	if err != nil {
// 		log.Printf("error when fetching data from database: %s\n", err.Error())
// 		errMsg = message.ErrorMessage{
// 			Error: err,
// 			Type:  "INTERNAL_CONNECTION_PROBLEM",
// 		}
// 		return result, errMsg
// 	}

// 	return result, errMsg
// }

// func (c *CommentUsecaseImpl) DeleteCommentSvc(ctx context.Context, commentId uint64) (errMsg message.ErrorMessage) {
// 	log.Printf("%T - DeleteCommentSvc is invoked\n", c)
// 	defer log.Printf("%T - DeleteCommentSvc executed\n", c)

// 	log.Println("calling delete comment repo")
// 	err := c.commentRepo.DeleteComment(ctx, commentId)

// 	if err != nil {
// 		log.Printf("error when fetching data from database: %s\n", err.Error())
// 		errMsg = message.ErrorMessage{
// 			Error: err,
// 			Type:  "INTERNAL_CONNECTION_PROBLEM",
// 		}
// 		return errMsg
// 	}

// 	return errMsg
// }

func NewCommentUsecase(commentRepo comment.CommentRepo, photoUsecase photo.PhotoUsecase) comment.CommentUsecase {
	return &CommentUsecaseImpl{commentRepo: commentRepo, photoUsecase: photoUsecase}
}
