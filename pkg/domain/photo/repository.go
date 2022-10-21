package photo

import "context"

type PhotoRepo interface {
	CreatePhoto(ctx context.Context, photo *Photo) (result Photo, err error)
	GetPhotosByUserId(ctx context.Context, userId uint64) (result []Photo, err error)
}
