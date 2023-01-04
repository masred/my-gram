package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/masred/my-gram/app/model/entity"
	"github.com/masred/my-gram/app/model/http/request"
	"github.com/masred/my-gram/app/model/http/response"
)

type photoService struct {
	PhotoRepository entity.PhotoRepository
	Validate        *validator.Validate
}

func NewPhotoService(photoRepository entity.PhotoRepository, validate *validator.Validate) entity.PhotoService {
	return &photoService{PhotoRepository: photoRepository, Validate: validate}
}

func (photoService *photoService) Create(payload request.PhotoCreate) (res response.PhotoCreate, err error) {

	if err = photoService.Validate.Struct(payload); err != nil {
		return
	}

	photo := entity.Photo{
		ID:       uuid.New(),
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoUrl: payload.PhotoUrl,
		UserID:   payload.UserID,
	}

	if err = photoService.PhotoRepository.Create(&photo); err != nil {
		return
	}

	res = response.PhotoCreate{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	}

	return
}

func (photoService *photoService) GetAll() (res []response.PhotoGetAll, err error) {
	var photos []entity.Photo
	photos, err = photoService.PhotoRepository.GetAll()
	if err != nil {
		return
	}

	for _, photo := range photos {
		res = append(res, response.PhotoGetAll{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: response.PhotoUserGetAll{
				Email:    photo.User.Email,
				Username: photo.User.Username,
			},
		})
	}

	return
}

func (photoService *photoService) GetOne(id uuid.UUID) (res response.PhotoGetOne, err error) {
	var photo entity.Photo
	photo, err = photoService.PhotoRepository.GetOne(id)
	if err != nil {
		return
	}

	res = response.PhotoGetOne{
		UserID: photo.UserID,
	}

	return
}

func (photoService *photoService) Update(id uuid.UUID, payload request.PhotoUpdate) (res response.PhotoUpdate, err error) {
	if err = photoService.Validate.Struct(payload); err != nil {
		return
	}

	photo := entity.Photo{
		ID:       id,
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoUrl: payload.PhotoUrl,
	}

	if err = photoService.PhotoRepository.Update(&photo); err != nil {
		return
	}

	res = response.PhotoUpdate{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}

	return
}

func (photoService *photoService) Delete(id uuid.UUID) (res response.PhotoDelete, err error) {
	if err = photoService.PhotoRepository.Delete(id); err != nil {
		return
	}

	res = response.PhotoDelete{
		Message: "photo successfully deleted",
	}

	return
}
