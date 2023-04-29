package handler

import (
	"fmt"
	"net/http"

	h "gojob/handler"
	"gojob/schemas"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v0

// @Summary Delete opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /openings [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		h.SendError(
			ctx,
			http.StatusBadRequest,
			"Id (string)")
		return
	}

	opening := schemas.Opening{}

	// Find Opening
	if err := h.DB.First(&opening, id).Error; err != nil {
		h.SendError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	// Delete Opening
	if err := h.DB.Delete(&opening).Error; err != nil {
		h.SendError(
			ctx,
			http.StatusInternalServerError,
			fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	h.SendSuccess(ctx, "delete-opening", opening)
}
