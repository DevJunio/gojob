package handler

import (
	"gojob/handler"
	"gojob/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := handler.UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("validation error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id"); if id == "" {
		handler.SendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	opening := schemas.Opening{}

	if err := handler.DB.First(&opening, id).Error; err != nil {
		handler.SendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	// Update opening
	switch {
	case request.Role     != ""  : opening.Role = request.Role
	case request.Company  != ""  : opening.Company = request.Company
	case request.Location != ""  : opening.Location = request.Location
	case request.Remote   != nil : opening.Remote = *request.Remote
	case request.Link     != ""  : opening.Link = request.Link
	case request.Salary   >  00  : opening.Salary = request.Salary }

	// Save opening
	if err := handler.DB.Save(&opening).Error; err != nil {
		handler.Logger.Errorf("error updating opening: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	handler.SendSuccess(ctx, "update-opening", opening)
}
