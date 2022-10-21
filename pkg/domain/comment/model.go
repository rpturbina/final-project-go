package comment

import (
	"github.com/rpturbina/final-project-go/pkg/domain/gormmodel"
)

type Comment struct {
	gormmodel.GormModel
	// TODO: give field tags
	UserID  uint64 `json:"user_id"`
	PhotoID uint64 `json:"photo_id"`
	Message string `gorm:"not null" json:"message" valid:"required~Comment message is required"`
}
