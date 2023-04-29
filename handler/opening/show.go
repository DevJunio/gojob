package handler

import (
	"net/http"

	"gojob/schemas"
	h "gojob/handler"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening with id
// @Tags Openings
// @Accept json
// @Produce json
// @Param id path int true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /openings/{id} [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		h.SendError(
			ctx,
			http.StatusBadRequest,
			h.FieldRequired("queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := h.DB.First(&opening, id).Error; err != nil {
		h.SendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	h.SendSuccess(ctx, "show-opening", opening)
}
