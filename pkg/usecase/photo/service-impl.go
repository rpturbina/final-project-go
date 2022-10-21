package photo

import (
	"context"
	"log"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/rpturbina/final-project-go/pkg/domain/message"
	"github.com/rpturbina/final-project-go/pkg/domain/photo"
)

type PhotoUsecaseImpl struct {
	photoRepo photo.PhotoRepo
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

func NewPhotoUsecase(photoRepo photo.PhotoRepo) photo.PhotoUsecase {
	return &PhotoUsecaseImpl{photoRepo: photoRepo}
}
