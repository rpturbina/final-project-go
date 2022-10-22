package comment

import (
	"context"
	"log"

	"github.com/rpturbina/final-project-go/config/postgres"
	"github.com/rpturbina/final-project-go/pkg/domain/comment"
	"gorm.io/gorm/clause"
)

type CommentRepoImpl struct {
	pgCln postgres.PostgresClient
}

func (c *CommentRepoImpl) CreateComment(ctx context.Context, inputComment *comment.Comment) (result comment.Comment, err error) {
	log.Printf("%T - CreateComment is invoked\n", c)
	defer log.Printf("%T - CreateComment executed\n", c)

	db := c.pgCln.GetClient()

	err = db.Model(&result).Create(&inputComment).Error

	if err != nil {
		log.Printf("error when creating comment for photo id %v\n", inputComment.PhotoID)
	}

	result = *inputComment

	return result, err
}

func (c *CommentRepoImpl) GetComments(ctx context.Context, userId uint64) (result []comment.Comment, err error) {
	log.Printf("%T - GetComments is invoked\n", c)
	defer log.Printf("%T - GetComments executed\n", c)

	db := c.pgCln.GetClient()

	err = db.Model(&comment.Comment{}).Where("user_id = ?", userId).Find(&result).Error

	if err != nil {
		log.Printf("error when getting photos by user id %v\n", userId)
	}

	return result, err
}

func (c *CommentRepoImpl) GetCommentById(ctx context.Context, commentId uint64) (result comment.Comment, err error) {
	log.Printf("%T - GetCommentById is invoked\n", c)
	defer log.Printf("%T - GetCommentById executed\n", c)

	db := c.pgCln.GetClient()

	err = db.Table("comments").Where("id = ?", commentId).Select("id", "message", "user_id", "photo_id").Find(&result).Error

	if err != nil {
		log.Printf("error when getting comment by id %v\n", commentId)
	}

	return result, err
}

func (c *CommentRepoImpl) UpdateComment(ctx context.Context, inputMessage string) (result comment.Comment, err error) {
	log.Printf("%T - UpdateComment is invoked\n", c)
	defer log.Printf("%T - UpdateComment executed\n", c)

	commentId := ctx.Value("commentId").(uint64)

	db := c.pgCln.GetClient()

	err = db.Model(&result).Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}, {Name: "photo_id"}, {Name: "message"}, {Name: "user_id"}, {Name: "updated_at"}}}).Where("id = ?", commentId).Updates(comment.Comment{Message: inputMessage}).Error

	if err != nil {
		log.Printf("error when updating comment by id %v\n", commentId)
	}

	return result, err
}

func (c *CommentRepoImpl) DeleteComment(ctx context.Context, commentId uint64) (err error) {
	log.Printf("%T - DeleteComment is invoked\n", c)
	defer log.Printf("%T - DeleteComment executed\n", c)

	db := c.pgCln.GetClient()

	err = db.Where("id = ?", commentId).Delete(&comment.Comment{}).Error

	if err != nil {
		log.Printf("error when deleting comment by id %v \n", commentId)
	}

	return err
}

func NewCommentRepo(pgCln postgres.PostgresClient) comment.CommentRepo {
	return &CommentRepoImpl{pgCln: pgCln}
}
