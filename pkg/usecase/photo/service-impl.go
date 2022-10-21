package photo

import (
	"context"
	"log"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/rpturbina/final-project-go/pkg/domain/message"
	"github.com/rpturbina/final-project-go/pkg/domain/photo"
	"github.com/rpturbina/final-project-go/pkg/domain/user"
)

type PhotoUsecaseImpl struct {
	photoRepo   photo.PhotoRepo
	userUsecase user.UserUsecase
}

func (p *PhotoUsecaseImpl) CreatePhotoSvc(ctx context.Context, photo photo.Photo) (result photo.Photo, errMsg message.ErrorMessage) {
	log.Printf("%T - CreatePhotoSvc is invoked\n", p)
	defer log.Printf("%T - CreatePhotoSvc executed\n", p)

	if isValid, err := govalidator.ValidateStruct(photo); !isValid {
		switch err.Error() {
		case "photo title is required":
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "PHOTO_TITLE_IS_EMPTY",
			}
			return result, errMsg
		case "photo url is required":
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "PHOTO_URL_IS_EMPTY",
			}
			return result, errMsg
		case "invalid url format":
			errMsg := message.ErrorMessage{
				Error: err,
				Type:  "INVALID_PHOTO_URL_FORMAT",
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

	photo.UserID = userId

	log.Println("calling create photo repo")
	result, err := p.photoRepo.CreatePhoto(ctx, &photo)

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

func (p *PhotoUsecaseImpl) GetPhotosByUserIdSvc(ctx context.Context, userId uint64) (result []photo.Photo, errMsg message.ErrorMessage) {
	log.Printf("%T - GetPhotosByUserIdSvc is invoked\n", p)
	defer log.Printf("%T - GetPhotosByUserIdSvc executed\n", p)

	log.Println("calling get user by id repo")
	checkUserId, errMsg := p.userUsecase.GetUserByIdSvc(ctx, userId)

	if checkUserId.ID <= 0 {
		return result, errMsg
	}

	log.Println("calling get photo by userid repo")
	result, err := p.photoRepo.GetPhotosByUserId(ctx, userId)

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

func NewPhotoUsecase(photoRepo photo.PhotoRepo, userUsecase user.UserUsecase) photo.PhotoUsecase {
	return &PhotoUsecaseImpl{photoRepo: photoRepo, userUsecase: userUsecase}
}