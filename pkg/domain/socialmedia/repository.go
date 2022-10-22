package socialmedia

import (
	"context"
)

type SocialMediaRepo interface {
	CreateSocialMedia(ctx context.Context, socialMedia *SocialMedia) (result SocialMedia, err error)
	GetSocialMedias(ctx context.Context, userId uint64) (result []SocialMedia, err error)
	GetSocialMediaById(ctx context.Context, socmedId uint64) (result SocialMedia, err error)
	UpdateSocialMedia(ctx context.Context, socialMedia SocialMedia) (result SocialMedia, err error)
	DeleteSocialMedia(ctx context.Context, socmedId uint64) (err error)
}
