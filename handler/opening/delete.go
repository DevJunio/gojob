package handler

import (
	"fmt"
	"net/http"

	h "gojob/handler"
	"gojob/schemas"

	"github.com/gin-gonic/gin"
)

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
