package repository

import (
	"errors"
	"github.com/devjunio/gojob/model"

	"github.com/gin-gonic/gin"
)

type OpeningRepository struct {
	request model.OpeningRequest

	ctx *gin.Context
	id  string
}

func (o *OpeningRepository) Get(op *model.Opening, id string) (*model.Opening, error) {
	result := db.First(&op, id)
	if err := result.Error; err != nil {
		return nil, err
	}

	return op, nil
}

func (o *OpeningRepository) List(openings []*model.Opening) ([]*model.Opening, error) {
	result := db.Find(&openings)
	if err := result.Error; err != nil {
		return nil, err
	}
	return openings, nil
}

func (o *OpeningRepository) Create(opening *model.Opening) (*model.Opening, error) {
	res := db.Create(&opening)
	if err := res.Error; err != nil {
		return nil, err
	}

	return opening, nil
}

func (o *OpeningRepository) Delete(opening *model.Opening, id string) (*model.OpeningResponse, error) {
	err := db.First(&opening, id).Error
	if err != nil {
		return nil, err
	}

	// Delete Opening
	if err = db.Delete(&opening, id).Error; err != nil {
		return nil, err
	}

	return opening.ToResponse(), nil
}

func (self *OpeningRepository) Update(opening *model.Opening, id string) (*model.OpeningResponse, []error) {

	if err := db.First(&opening, id).Error; err != nil {
		return nil, []error{errors.New("first"), err}
	}

	if err := db.Save(&opening).Error; err != nil {
		return nil, []error{errors.New("save"), err}
	}

	return opening.ToResponse(), nil
}
