package service

import (
	"net/http"

	"github.com/devjunio/gojob/config"
	"github.com/devjunio/gojob/model"
	"github.com/devjunio/gojob/repository"

	"github.com/gin-gonic/gin"
)

func InitService() *OpeningService {
	logger = config.SetLogger("service")
	return &OpeningService{}
}

var openingRepository *repository.OpeningRepository

type OpeningService struct {
	request *model.OpeningRequest
}

var (
	logger *config.Logger
)

// Get
//
//	@Summary		Get Openings
//	@Description	Get first opening by ID
//	@Tags			Openings
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string				true	"Opening identification"
//	@Success		200	{object}	model.OpeningResponse
//	@Failure		500	{object}	router.ErrorResponse
//	@Router			/openings/:id [get]
func (o *OpeningService) Get(ctx *gin.Context) {
	var opening *model.Opening

	id := ctx.Param("id")
	if id == "" {
		logger.Error("User not found")
		sendError(ctx, http.StatusNotFound, "User not found")
		return
	}

	response, err := openingRepository.Get(opening, id)
	if err != nil {
		logger.Errorf("error getting item (opening): %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
	}

	sendSuccess(ctx, response)
}

// List
//
//	@Summary		List Openings
//	@Description	List all job opening
//	@Tags			Openings
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.OpeningResponse
//	@Failure		500	{object}	router.ErrorResponse
//	@Router			/openings [get]
func (o *OpeningService) List(ctx *gin.Context) {
	var request []*model.Opening
	openings, err := openingRepository.List(request)
	if err != nil {
		logger.Errorf("error getting item (opening): %v", err.Error())
		sendError(ctx, http.StatusNotFound, "opening not found")
	}

	logger.Info("List retrieved successfully")
	sendSuccess(ctx, openings)
}

// Create
//
//	@Summary		Create opening
//	@Description	Create a new job opening
//	@Tags			Openings
//	@Accept			json
//	@Produce		json
//	@Param			opening	body		model.OpeningRequest	true	"Opening data to create"
//	@Success		200		{object}	model.OpeningResponse
//	@Failure		400		{object}	router.ErrorResponse
//	@Failure		404		{object}	router.ErrorResponse
//	@Router			/openings [post]
func (o *OpeningService) Create(ctx *gin.Context) {
	request := model.OpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("Binding request JSON error: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := request.ValidateCreation(); err != nil {
		logger.Errorf("Validating request fields error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response, err := openingRepository.Create(request.ToOpening())
	if err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError,
			"error while creating Opening on database")
		return
	}

	sendSuccess(ctx, response)
}

// Delete
//
//	@Summary		Delete opening
//	@Description	Delete a new job opening
//	@Tags			Openings
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"Opening identification"
//	@Success		200	{object}	model.OpeningResponse
//	@Failure		400	{object}	router.ErrorResponse
//	@Failure		404	{object}	router.ErrorResponse
//	@Router			/openings [delete]
func (o *OpeningService) Delete(ctx *gin.Context) {
	opening := model.Opening{}

	id := getID(ctx)
	if id == "" {
		return
	}

	response, err := openingRepository.Delete(&opening, id)
	if err != nil {
		logger.Errorf("error while deleting opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error on deleting")
		return
	}

	sendSuccess(ctx, response)
}

// Update
//
//	@Summary		Update opening
//	@Description	Update a job opening
//	@Tags			Openings
//	@Accept			json
//	@Produce		json
//	@Param			id		query		string					true	"Opening Identification"
//	@Param			opening	body		model.OpeningRequest	true	"Opening data to Update"
//	@Success		200		{object}	model.OpeningResponse
//	@Failure		400		{object}	router.ErrorResponse
//	@Failure		404		{object}	router.ErrorResponse
//	@Failure		500		{object}	router.ErrorResponse
//	@Router			/openings [patch]
func (o *OpeningService) Update(ctx *gin.Context) {
	var request *model.OpeningRequest
	id := getID(ctx)
	if id == "" {
		return
	}

	if err := request.ValidateUpdate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "missing field")
		return
	}

	res, err := openingRepository.Update(request.ToOpening(), id)
	if err != nil {
		switch err[0].Error() {
		case "first":
			logger.Errorf("Getting first item: %v", err[1].Error())
			sendError(ctx, http.StatusNotFound, "getting item")

		case "save":
			logger.Errorf("Updating error: %v", err[1].Error())
			sendError(ctx, http.StatusInternalServerError, "error on updating opening")

		default:
			logger.Errorf("Unknown error: %v", err[1].Error())
			sendError(ctx, http.StatusForbidden, "unknown")
		}
	}

	// err: valid
	sendSuccess(ctx, res)
}
