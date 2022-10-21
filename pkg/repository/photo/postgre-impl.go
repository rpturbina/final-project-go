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

func NewPhotoRepo(pgCln postgres.PostgresClient) photo.PhotoRepo {
	return &PhotoRepoImpl{pgCln: pgCln}
}
