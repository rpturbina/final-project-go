package photo

import (
	"github.com/asaskevich/govalidator"
	"github.com/rpturbina/final-project-go/pkg/domain/comment"
	"github.com/rpturbina/final-project-go/pkg/domain/gormmodel"
	"gorm.io/gorm"
)

type Photo struct {
	gormmodel.GormModel
	// TODO: give field tags
	Title    string            `gorm:"not null" json:"title"`
	Caption  string            `json:"caption"`
	Url      string            `gorm:"not null" json:"url" valid:"required~Photo url is required"`
	UserID   uint64            `json:"user_id"`
	Comments []comment.Comment `json:"comments"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
