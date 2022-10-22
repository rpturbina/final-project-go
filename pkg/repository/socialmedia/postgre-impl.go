package socialmedia

import (
	"context"
	"log"

	"github.com/rpturbina/final-project-go/config/postgres"
	"github.com/rpturbina/final-project-go/pkg/domain/socialmedia"
)

type SocialMediaRepoImpl struct {
	pgCln postgres.PostgresClient
}

func (s *SocialMediaRepoImpl) CreateSocialMedia(ctx context.Context, inputSocialMedia *socialmedia.SocialMedia) (result socialmedia.SocialMedia, err error) {
	log.Printf("%T - CreateSocialMedia is invoked\n", s)
	defer log.Printf("%T - CreateSocialMedia executed\n", s)

	db := s.pgCln.GetClient()

	err = db.Model(&result).Create(&inputSocialMedia).Error

	if err != nil {
		log.Printf("error when creating socialMedia for photo id %v\n", inputSocialMedia)
	}

	result = *inputSocialMedia

	return result, err
}

// func (s *SocialMediaRepoImpl) GetSocialMedias(ctx context.Context, userId uint64) (result []socialMedia.SocialMedia, err error) {
// 	log.Printf("%T - GetSocialMedias is invoked\n", s)
// 	defer log.Printf("%T - GetSocialMedias executed\n", s)

// 	db := s.pgCln.GetClient()

// 	err = db.Model(&socialMedia.SocialMedia{}).Where("user_id = ?", userId).Find(&result).Error

// 	if err != nil {
// 		log.Printf("error when getting photos by user id %v\n", userId)
// 	}

// 	return result, err
// }

// func (s *SocialMediaRepoImpl) GetSocialMediaById(ctx context.Context, socialMediaId uint64) (result socialMedia.SocialMedia, err error) {
// 	log.Printf("%T - GetSocialMediaById is invoked\n", s)
// 	defer log.Printf("%T - GetSocialMediaById executed\n", s)

// 	db := s.pgCln.GetClient()

// 	err = db.Table("socialMedias").Where("id = ?", socialMediaId).Select("id", "message", "user_id", "photo_id").Find(&result).Error

// 	if err != nil {
// 		log.Printf("error when getting socialMedia by id %v\n", socialMediaId)
// 	}

// 	return result, err
// }

// func (s *SocialMediaRepoImpl) UpdateSocialMedia(ctx context.Context, inputMessage string) (result socialMedia.SocialMedia, err error) {
// 	log.Printf("%T - UpdateSocialMedia is invoked\n", s)
// 	defer log.Printf("%T - UpdateSocialMedia executed\n", s)

// 	socialMediaId := ctx.Value("socialMediaId").(uint64)

// 	db := s.pgCln.GetClient()

// 	err = db.Model(&result).Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}, {Name: "photo_id"}, {Name: "message"}, {Name: "user_id"}, {Name: "updated_at"}}}).Where("id = ?", socialMediaId).Updates(socialMedia.SocialMedia{Message: inputMessage}).Error

// 	if err != nil {
// 		log.Printf("error when updating socialMedia by id %v\n", socialMediaId)
// 	}

// 	return result, err
// }

// func (s *SocialMediaRepoImpl) DeleteSocialMedia(ctx context.Context, socialMediaId uint64) (err error) {
// 	log.Printf("%T - DeleteSocialMedia is invoked\n", s)
// 	defer log.Printf("%T - DeleteSocialMedia executed\n", s)

// 	db := s.pgCln.GetClient()

// 	err = db.Where("id = ?", socialMediaId).Delete(&socialMedia.SocialMedia{}).Error

// 	if err != nil {
// 		log.Printf("error when deleting socialMedia by id %v \n", socialMediaId)
// 	}

// 	return err
// }

func NewSocialMediaRepo(pgCln postgres.PostgresClient) socialmedia.SocialMediaRepo {
	return &SocialMediaRepoImpl{pgCln: pgCln}
}
