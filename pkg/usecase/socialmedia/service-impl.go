package socialmedia

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/rpturbina/final-project-go/pkg/domain/message"
	"github.com/rpturbina/final-project-go/pkg/domain/socialmedia"
)

type SocialMediaUsecaseImpl struct {
	socialMediaRepo socialmedia.SocialMediaRepo
}

func (s *SocialMediaUsecaseImpl) CreateSocialMediaSvc(ctx context.Context, socialMedia socialmedia.SocialMedia) (result socialmedia.SocialMedia, errMsg message.ErrorMessage) {
	log.Printf("%T - CreateSocialMediaSvc is invoked\n", s)
	defer log.Printf("%T - CreateSocialMediaSvc executed\n", s)

	if isValid, err := govalidator.ValidateStruct(socialMedia); !isValid {
		switch err.Error() {
		case "social media name is required":
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "SOCIAL_MEDIA_NAME_IS_EMPTY",
			}
			return result, errMsg
		case "social media url is required":
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "SOCIAL_MEDIA_URL_IS_EMPTY",
			}
			return result, errMsg
		case "invalid url format":
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "INVALID_URL_FORMAT",
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

	stringUserId := ctx.Value("user").(string)

	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

	socialMedia.UserID = userId

	log.Println("calling create socialMedia repo")
	result, err := s.socialMediaRepo.CreateSocialMedia(ctx, &socialMedia)

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

func (s *SocialMediaUsecaseImpl) GetSocialMediasSvc(ctx context.Context) (result []socialmedia.SocialMedia, errMsg message.ErrorMessage) {
	log.Printf("%T - GetSocialMediasSvc is invoked\n", s)
	defer log.Printf("%T - GetSocialMediasSvc executed\n", s)

	stringUserId := ctx.Value("user").(string)

	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

	log.Println("calling get social media by user id repo")
	result, err := s.socialMediaRepo.GetSocialMedias(ctx, userId)

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

func (s *SocialMediaUsecaseImpl) GetSocialMediaByIdSvc(ctx context.Context, socmedId uint64) (result socialmedia.SocialMedia, errMsg message.ErrorMessage) {
	log.Printf("%T - GetSocialMediaByIdSvc is invoked\n", s)
	defer log.Printf("%T - GetSocialMediaByIdSvc executed\n", s)

	log.Println("calling get social media by id repo")
	result, err := s.socialMediaRepo.GetSocialMediaById(ctx, socmedId)

	if result.ID <= 0 {
		log.Printf("social media with id %v not found", socmedId)

		err = fmt.Errorf("social media with id %v not found", socmedId)
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "SOCIAL_MEDIA_NOT_FOUND",
		}
		return result, errMsg
	}

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

func (s *SocialMediaUsecaseImpl) UpdateSocialMediaSvc(ctx context.Context, inputSocialMedia socialmedia.SocialMedia) (result socialmedia.SocialMedia, errMsg message.ErrorMessage) {
	log.Printf("%T - UpdateSocialMediaSvc is invoked\n", s)
	defer log.Printf("%T - UpdateSocialMediaSvc executed\n", s)

	if isValid, err := govalidator.ValidateStruct(inputSocialMedia); !isValid {
		switch err.Error() {
		case "social media name is required":
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "SOCIAL_MEDIA_NAME_IS_EMPTY",
			}
			return result, errMsg
		case "social media url is required":
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "SOCIAL_MEDIA_URL_IS_EMPTY",
			}
			return result, errMsg
		case "invalid url format":
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "INVALID_URL_FORMAT",
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

	result, err := s.socialMediaRepo.UpdateSocialMedia(ctx, inputSocialMedia)

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

func (s *SocialMediaUsecaseImpl) DeleteSocialMediaSvc(ctx context.Context, socmedId uint64) (errMsg message.ErrorMessage) {
	log.Printf("%T - DeleteSocialMediaSvc is invoked\n", s)
	defer log.Printf("%T - DeleteSocialMediaSvc executed\n", s)

	log.Println("calling delete socialMedia repo")
	err := s.socialMediaRepo.DeleteSocialMedia(ctx, socmedId)

	if err != nil {
		log.Printf("error when fetching data from database: %s\n", err.Error())
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "INTERNAL_CONNECTION_PROBLEM",
		}
		return errMsg
	}

	return errMsg
}

func NewSocialMediaUsecase(socialMediaRepo socialmedia.SocialMediaRepo) socialmedia.SocialMediaUsecase {
	return &SocialMediaUsecaseImpl{socialMediaRepo: socialMediaRepo}
}
