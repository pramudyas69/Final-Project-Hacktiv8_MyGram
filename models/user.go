package models

import (
	"MyGram/utils"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Username  string `json:"username" gorm:"unique;not null" form:"username" valid:"required~Username is required"`
	Email     string `json:"email" gorm:"unique;not null" form:"email" valid:"required~Email is required, email~Email is invalid"`
	Password  string `json:"password" gorm:"not null" form:"password" valid:"required~Password is required, minstringlength(6)~Password must be at least 6 characters"`
	Age       int    `json:"age" gorm:"not null" form:"age" valid:"required~Age is required, range(8|100)~Age must be at least 8"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		return errCreate
	}

	u.Password = utils.HashPassword(u.Password)
	u.CreatedAt = time.Now().Local()
	return nil
}

func (u *User) BeforeUpdate(db *gorm.DB) error {
	u.UpdatedAt = time.Now().Local()
	return nil
}
