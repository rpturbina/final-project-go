package photo

import (
	"context"
	"log"

	"github.com/rpturbina/final-project-go/config/postgres"
	"github.com/rpturbina/final-project-go/pkg/domain/photo"
)

type PhotoRepoImpl struct {
	pgCln postgres.PostgresClient
}

func (p *PhotoRepoImpl) CreatePhoto(ctx context.Context, inputPhoto *photo.Photo) (result photo.Photo, err error) {
	log.Printf("%T - CreatePhoto is invoked\n", p)
	defer log.Printf("%T - CreatePhoto executed\n", p)

	db := p.pgCln.GetClient()

	err = db.Model(&result).Create(&inputPhoto).Error

	if err != nil {
		log.Printf("error when creating photo %v\n", inputPhoto.ID)
	}

	result = *inputPhoto

	return result, err
}

func (p *PhotoRepoImpl) GetPhotosByUserId(ctx context.Context, userId uint64) (result []photo.Photo, err error) {
	log.Printf("%T - GetPhotosByUserId is invoked\n", p)
	defer log.Printf("%T - GetPhotosByUserId executed\n", p)

	db := p.pgCln.GetClient()

	err = db.Table("photos").Where("user_id = ?", userId).Select("id", "created_at", "updated_at", "title", "caption", "url", "user_id").Order("id").Preload("Comments").Find(&result).Error

	if err != nil {
		log.Printf("error when getting photos by user id %v\n", userId)
	}

	return result, err
}

func (p *PhotoRepoImpl) GetPhotoById(ctx context.Context, photoId uint64) (result photo.Photo, err error) {
	log.Printf("%T - GetPhotoById is invoked\n", p)
	defer log.Printf("%T - GetPhotoById executed\n", p)

	db := p.pgCln.GetClient()

	err = db.Table("photos").Where("id = ?", photoId).Select("id", "title", "caption", "url", "user_id").Preload("Comments").Find(&result).Error

	if err != nil {
		log.Printf("error when getting photo by id %v\n", photoId)
	}

	return result, err
}

func NewPhotoRepo(pgCln postgres.PostgresClient) photo.PhotoRepo {
	return &PhotoRepoImpl{pgCln: pgCln}
}
