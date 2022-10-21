package photo

import (
	"context"
	"log"

	"github.com/rpturbina/final-project-go/config/postgres"
	"github.com/rpturbina/final-project-go/pkg/domain/comment"
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

// func (c *CommentRepoImpl) GetPhotosByUserId(ctx context.Context, userId uint64) (result []photo.Photo, err error) {
// 	log.Printf("%T - GetPhotosByUserId is invoked\n", c)
// 	defer log.Printf("%T - GetPhotosByUserId executed\n", c)

// 	db := c.pgCln.GetClient()

// 	err = db.Table("photos").Where("user_id = ?", userId).Select("id", "created_at", "updated_at", "title", "caption", "url", "user_id").Order("id").Preload("Comments").Find(&result).Error

// 	if err != nil {
// 		log.Printf("error when getting photos by user id %v\n", userId)
// 	}

// 	return result, err
// }

// func (c *CommentRepoImpl) GetPhotoById(ctx context.Context, photoId uint64) (result photo.Photo, err error) {
// 	log.Printf("%T - GetPhotoById is invoked\n", c)
// 	defer log.Printf("%T - GetPhotoById executed\n", c)

// 	db := c.pgCln.GetClient()

// 	err = db.Table("photos").Where("id = ?", photoId).Select("id", "title", "caption", "url", "user_id").Preload("Comments").Find(&result).Error

// 	if err != nil {
// 		log.Printf("error when getting photo by id %v\n", photoId)
// 	}

// 	return result, err
// }

// func (c *CommentRepoImpl) UpdatePhoto(ctx context.Context, title string, caption string, url string) (result photo.Photo, err error) {
// 	log.Printf("%T - UpdatePhoto is invoked\n", c)
// 	defer log.Printf("%T - UpdatePhoto executed\n", c)

// 	photoId := ctx.Value("photoId").(uint64)

// 	db := c.pgCln.GetClient()

// 	err = db.Model(&result).Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}, {Name: "title"}, {Name: "caption"}, {Name: "url"}, {Name: "user_id"}, {Name: "updated_at"}}}).Where("id = ?", photoId).Updates(photo.Photo{Title: title, Caption: caption, Url: url}).Error

// 	if err != nil {
// 		log.Printf("error when updating photo by id %v\n", photoId)
// 	}

// 	return result, err
// }

// func (c *CommentRepoImpl) DeletePhoto(ctx context.Context, photoId uint64) (err error) {
// 	log.Printf("%T - DeletePhoto is invoked\n", c)
// 	defer log.Printf("%T - DeletePhoto executed\n", c)

// 	db := c.pgCln.GetClient()

// 	err = db.Where("id = ?", photoId).Delete(&photo.Photo{}).Error

// 	if err != nil {
// 		log.Printf("error when deleting photo by id %v \n", photoId)
// 	}

// 	return err
// }

func NewCommentRepo(pgCln postgres.PostgresClient) comment.CommentRepo {
	return &CommentRepoImpl{pgCln: pgCln}
}
