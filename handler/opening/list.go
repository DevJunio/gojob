package handler

import (
	"net/http"

	"gojob/schemas"
	h "gojob/handler"

	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := h.DB.Find(&openings).Error; err != nil {
		h.SendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	h.SendSuccess(ctx, "list-openings", openings)
}
