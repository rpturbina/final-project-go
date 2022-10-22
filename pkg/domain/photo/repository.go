package photo

import (
	"context"
)

type PhotoRepo interface {
	CreatePhoto(ctx context.Context, photo *Photo) (result Photo, err error)
	GetPhotosByUserId(ctx context.Context, userId uint64) (result []Photo, err error)
	GetPhotoById(ctx context.Context, photoId uint64) (result Photo, err error)
	UpdatePhoto(ctx context.Context, title string, caption string, url string) (result Photo, err error)
	DeletePhoto(ctx context.Context, photoId uint64) (err error)
}
