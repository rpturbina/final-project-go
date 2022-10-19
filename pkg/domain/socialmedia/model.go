package socialmedia

import (
	"github.com/asaskevich/govalidator"
	"github.com/rpturbina/final-project-go/pkg/domain/gormmodel"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gormmodel.GormModel
	Name   string `gorm:"not null" json:"name" valid:"required~Social media name is required"`
	URL    string `gorm:"not null" json:"url" valid:"required~Social media URL is required"`
	UserID uint64 `json:"user_id"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
