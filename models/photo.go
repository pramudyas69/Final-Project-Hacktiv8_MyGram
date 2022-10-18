package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"time"
)

type Photo struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Title     string `json:"title" gorm:"not null" form:"title" valid:"required~Title is required"`
	Caption   string `json:"caption" form:"caption"`
	PhotoUrl  string `json:"photo_url" gorm:"not null" form:"photo_url" valid:"required~PhotoUrl is required"`
	UserId    uint   `json:"user_id" form:"user_id"`
	User      *User  `json:"User"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	p.CreatedAt = time.Now().Local()

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	p.UpdatedAt = time.Now().Local()

	err = nil
	return
}
