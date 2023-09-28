package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"guacamole_api/internal/model"
	"net/http"
)

func (c Controller) Users(ctx echo.Context) error {
	response, responseError := http.Get(fmt.Sprintf("http://%s/guacamole/api/session/data/%s/users?token=%s", c.Url, ctx.Param(":source"), ctx.QueryParam("token")))
	defer response.Body.Close()

	if responseError != nil {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusOK,
			Message:    "error response users credentials",
		})
	}
	return ctx.JSON(http.StatusOK, model.SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    "success response authentication",
		Data:       "responseAuthTokens",
	})
}
