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

// Get Job Opening by ID
//
// Retrieves detailed information about the first job opening using its unique identifier.
// This endpoint offers a comprehensive view of a specific employment opportunity.
//
// @Summary Retrieve Job Opening Details by ID
// @Description Fetch detailed information about the first job opening based on its unique identifier.
//
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

// List Job Openings
//
// Retrieves a comprehensive list of all available job openings in the system.
// This endpoint provides an overview of current employment opportunities.
//
// @Summary Retrieve Job Opening List
// @Description Retrieve a complete list of job openings stored in the system.
//
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

// Create Job Opening
//
// Initiates the creation of a new job opening resource by processing a JSON request.
// This endpoint enables the addition of fresh employment opportunities to the system.
//
//	@Summary		Initiate Job Opening Creation
//	@Description	Add a new job opening to the system using a JSON request.
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

// Delete Job Opening
//
// Removes a job opening from the system using its designated identifier.
// This endpoint facilitates the removal of a specific employment opportunity.
//
// @Summary Remove Job Opening
// @Description Delete a job opening resource based on its unique identifier.
//
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

// Update Job Opening
//
// Modifies an existing job opening in the system using appropriate changes provided in the request.
// This endpoint facilitates the adjustment of details for a specific employment opportunity.
//
// @Summary Modify Job Opening
// @Description Update an existing job opening with the provided changes.
//
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
