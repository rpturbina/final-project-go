package socialmedia

import (
	"context"
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

// func (s *SocialMediaUsecaseImpl) GetSocialMediasSvc(ctx context.Context) (result []socialMedia.SocialMedia, errMsg message.ErrorMessage) {
// 	log.Printf("%T - GetSocialMediasSvc is invoked\n", s)
// 	defer log.Printf("%T - GetSocialMediasSvc executed\n", s)

// 	stringUserId := ctx.Value("user").(string)

// 	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

// 	log.Println("calling get socialMedia by userid repo")
// 	result, err := s.socialMediaRepo.GetSocialMedias(ctx, userId)

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

// func (s *SocialMediaUsecaseImpl) GetSocialMediaByIdSvc(ctx context.Context, socialMediaId uint64) (result socialMedia.SocialMedia, errMsg message.ErrorMessage) {
// 	log.Printf("%T - GetSocialMediaByIdSvc is invoked\n", s)
// 	defer log.Printf("%T - GetSocialMediaByIdSvc executed\n", s)

// 	log.Println("calling get socialMedia by id repo")
// 	result, err := s.socialMediaRepo.GetSocialMediaById(ctx, socialMediaId)

// 	if result.ID <= 0 {
// 		log.Printf("socialMedia with id %v not found", socialMediaId)

// 		err = fmt.Errorf("socialMedia with id %v not found", socialMediaId)
// 		errMsg = message.ErrorMessage{
// 			Error: err,
// 			Type:  "COMMENT_NOT_FOUND",
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

// func (s *SocialMediaUsecaseImpl) UpdateSocialMediaSvc(ctx context.Context, inputMessage string) (result socialMedia.SocialMedia, errMsg message.ErrorMessage) {
// 	log.Printf("%T - UpdateSocialMediaSvc is invoked\n", s)
// 	defer log.Printf("%T - UpdateSocialMediaSvc executed\n", s)

// 	if inputMessage == "" {
// 		errMsg := message.ErrorMessage{
// 			Error: errors.New("socialMedia message is required"),
// 			Type:  "COMMENT_MESSAGE_IS_EMPTY",
// 		}
// 		return result, errMsg
// 	}

// 	result, err := s.socialMediaRepo.UpdateSocialMedia(ctx, inputMessage)

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

// func (s *SocialMediaUsecaseImpl) DeleteSocialMediaSvc(ctx context.Context, socialMediaId uint64) (errMsg message.ErrorMessage) {
// 	log.Printf("%T - DeleteSocialMediaSvc is invoked\n", s)
// 	defer log.Printf("%T - DeleteSocialMediaSvc executed\n", s)

// 	log.Println("calling delete socialMedia repo")
// 	err := s.socialMediaRepo.DeleteSocialMedia(ctx, socialMediaId)

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

func NewSocialMediaUsecase(socialMediaRepo socialmedia.SocialMediaRepo) socialmedia.SocialMediaUsecase {
	return &SocialMediaUsecaseImpl{socialMediaRepo: socialMediaRepo}
}
