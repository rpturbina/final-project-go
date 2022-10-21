package photo

import "context"

type PhotoRepo interface {
	CreatePhoto(ctx context.Context, photo *Photo) (result Photo, err error)
}
