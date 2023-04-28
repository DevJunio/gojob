package handler

import (
	"net/http"

	"gojob/schemas"
	h "gojob/handler"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
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
