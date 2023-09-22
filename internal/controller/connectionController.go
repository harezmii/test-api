package controller

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"guacamole_api/internal/model"
	"net/http"
)

func (c Controller) GetConnections(ctx echo.Context) error {
	response, responseError := http.Get(fmt.Sprintf("http://%s/guacamole/api/session/data/postgresql/connections?token=%s", c.Url, ctx.QueryParam("token")))
	defer response.Body.Close()

	if responseError != nil {
		return ctx.JSON(http.StatusOK, model.ErrorResponse{
			StatusCode: http.StatusOK,
			Message:    "error response get connections",
		})
	}
	var d interface{}
	decodeError := json.NewDecoder(response.Body).Decode(&d)
	if decodeError != nil {
		return ctx.JSON(http.StatusOK, model.ErrorResponse{
			StatusCode: http.StatusOK,
			Message:    "error response get connections decode",
		})
	}

	return ctx.JSON(http.StatusOK, model.SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    "success response get connections",
		Data:       d,
	})
}
