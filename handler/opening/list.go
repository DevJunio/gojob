package handler

import (
	"net/http"

	"gojob/schemas"
	h "gojob/handler"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v0

// @Summary List Openings
// @Description List all job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} handler.ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := h.DB.Find(&openings).Error; err != nil {
		h.SendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	h.SendSuccess(ctx, "list-openings", openings)
}
