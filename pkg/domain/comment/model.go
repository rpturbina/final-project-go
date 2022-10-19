package comment

import (
	"github.com/asaskevich/govalidator"
	"github.com/rpturbina/final-project-go/pkg/domain/gormmodel"
	"gorm.io/gorm"
)

type Comment struct {
	gormmodel.GormModel
	// TODO: give field tags
	UserID  uint64 `json:"user_id"`
	PhotoID uint64 `json:"photo_id"`
	Message string `gorm:"not null" json:"message" valid:"required~Comment message is required"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
