package comment

import (
	"github.com/rpturbina/final-project-go/pkg/domain/gormmodel"
)

type Comment struct {
	gormmodel.GormModel
	UserID  uint64 `json:"user_id"`
	PhotoID uint64 `json:"photo_id" valid:"required~photo id is required"`
	Message string `gorm:"not null" json:"message" valid:"required~comment message is required"`
}
