package repository

import (
	"github.com/google/uuid"
	"github.com/masred/my-gram/app/model/entity"
	"gorm.io/gorm"
)

type photoRepository struct {
	Database *gorm.DB
}

func NewPhotoRepository(database *gorm.DB) entity.PhotoRepository {
	return &photoRepository{Database: database}
}

func (photoRepository *photoRepository) Create(photo *entity.Photo) (err error) {
	if err = photoRepository.Database.Create(&photo).Error; err != nil {
		return
	}

	return
}

func (photoRepository *photoRepository) GetAll() (photos []entity.Photo, err error) {
	if err = photoRepository.Database.Preload("User").Find(&photos).Error; err != nil {
		return
	}

	return
}

func (photoRepository *photoRepository) GetOne(id uuid.UUID) (photo entity.Photo, err error) {
	if err = photoRepository.Database.First(&photo, id).Error; err != nil {
		return
	}

	return
}

func (photoRepository *photoRepository) Update(photo *entity.Photo) (err error) {
	newPhoto := entity.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
	}

	if err = photoRepository.Database.First(&photo, photo.ID).Error; err != nil {
		return
	}

	if err = photoRepository.Database.Model(&photo).Updates(newPhoto).Error; err != nil {
		return
	}

	return
}

func (photoRepository *photoRepository) Delete(id uuid.UUID) (err error) {
	var photo entity.Photo

	if err = photoRepository.Database.First(&photo, id).Error; err != nil {
		return
	}

	if err = photoRepository.Database.Delete(&photo).Error; err != nil {
		return
	}

	return
}
