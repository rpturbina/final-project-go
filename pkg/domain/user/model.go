package user

import (
	"encoding/json"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/rpturbina/final-project-go/helpers"
	"github.com/rpturbina/final-project-go/pkg/domain/comment"
	"github.com/rpturbina/final-project-go/pkg/domain/gormmodel"
	"github.com/rpturbina/final-project-go/pkg/domain/photo"
	"github.com/rpturbina/final-project-go/pkg/domain/socialmedia"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type customTime datatypes.Date

var _ json.Unmarshaler = &customTime{}

func (mt *customTime) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*mt = customTime(t)
	return nil
}

type User struct {
	gormmodel.GormModel
	Username string `gorm:"not null;uniqueIndex" json:"username" valid:"required~Your username is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null" json:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	// TODO: give validation the minimum age is above 8
	DOB    customTime    `gorm:"not null" json:"dob" valid:"required~Your dob is required"`
	Photos []photo.Photo `json:"photos"`
	// TODO: give comments and socialmedia field tags
	Comments     []comment.Comment         `json:"comments"`
	SocialMedias []socialmedia.SocialMedia `json:"social_medias"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
