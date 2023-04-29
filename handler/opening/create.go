package handler

import (
	"github.com/gin-gonic/gin"
	h "gojob/handler"
	"gojob/schemas"
	"net/http"
)

// @BasePath /api/v0

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := h.CreateOpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		h.Logger.Error("Binding request JSON error: %v", err.Error())
		h.SendError(ctx, http.StatusBadRequest, err.Error())

		return
	}

	if err := request.Validate(); err != nil {
		h.Logger.Errorf("Validating request fields error: %v", err.Error())
		h.SendError(ctx, http.StatusBadRequest, err.Error())

		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := h.DB.Create(&opening).Error; err != nil {
		h.Logger.Errorf("error creating opening: %v", err.Error())
		h.SendError(
			ctx,
			http.StatusInternalServerError,
			"error creating opening on database")

		return
	}

	h.SendSuccess(ctx, "creating-opening", &opening)
}
