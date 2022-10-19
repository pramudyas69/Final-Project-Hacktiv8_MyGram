package models

import (
	"MyGram/utils"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	Age       int       `json:"age" gorm:"not null" form:"age" valid:"required~Age is required, range(8|100)~Age must be at least 8"`
	Email     string    `json:"email" gorm:"unique;not null" form:"email" valid:"required~Email is required, email~Email is invalid"`
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `json:"username" gorm:"unique;not null" form:"username" valid:"required~Username is required"`
	Password  string    `json:"password,omitempty" gorm:"not null" form:"password" valid:"required~Password is required, minstringlength(6)~Password must be at least 6 characters"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
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
