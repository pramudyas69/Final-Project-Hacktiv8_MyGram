package create

import (
	"MyGram/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateRepositoryPhoto(input *models.Photo) (*models.Photo, string)
}

type repository struct {
	db *gorm.DB
}

func NewCreateRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateRepositoryPhoto(input *models.Photo) (*models.Photo, string) {
	var photos models.Photo
	db := r.db.Model(&photos)
	errorCode := make(chan string, 1)

	photos.Title = input.Title
	photos.Caption = input.Caption
	photos.PhotoUrl = input.PhotoUrl

	addNewPhoto := db.Debug().Create(&photos)
	db.Commit()

	if addNewPhoto.Error != nil {
		errorCode <- "CREATE_STUDENT_FAILED_403"
		return &photos, <-errorCode
	} else {
		errorCode <- "nil"
	}
	return &photos, <-errorCode
}
