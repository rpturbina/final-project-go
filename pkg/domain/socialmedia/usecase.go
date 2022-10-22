package socialmedia

import (
	"context"

	"github.com/rpturbina/final-project-go/pkg/domain/message"
)

type SocialMediaUsecase interface {
	CreateSocialMediaSvc(ctx context.Context, socialMedia SocialMedia) (result SocialMedia, errMsg message.ErrorMessage)
	// GetSocialMediasSvc(ctx context.Context) (result []SocialMedia, errMsg message.ErrorMessage)
	// GetSocialMediaByIdSvc(ctx context.Context, socialMediaId uint64) (result SocialMedia, errMsg message.ErrorMessage)
	// UpdateSocialMediaSvc(ctx context.Context, inputMessage string) (result SocialMedia, errMsg message.ErrorMessage)
	// DeleteSocialMediaSvc(ctx context.Context, socialMediaId uint64) (errMsg message.ErrorMessage)
}
