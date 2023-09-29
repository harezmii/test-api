package controller

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"guacamole_api/internal/model"
	"net/http"
	"net/url"
)

func (c Controller) AuthTokens(ctx echo.Context) error {

	if ctx.FormValue("username") == "" || ctx.FormValue("password") == "" {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "missing or incorrect information bind lo de",
		})
	}
	postBody := url.Values{
		"username": {ctx.FormValue("username")},
		"password": {ctx.FormValue("password")},
	}
	response, responseError := http.PostForm(fmt.Sprintf("http://%s/guacamole/api/tokens", c.Url), postBody)
	if responseError != nil {
		fmt.Println(responseError.Error())
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "missing or incorrect information up",
		})
	}
	defer response.Body.Close()

	responseAuthTokens := model.ResponseAuthModel{}
	json.NewDecoder(response.Body).Decode(&responseAuthTokens)
	fmt.Println(responseAuthTokens)
	return ctx.JSON(http.StatusOK, model.SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    "success response authentication",
		Data:       responseAuthTokens,
	})
}

func (c Controller) DeleteTokens(ctx echo.Context) error {
	token := ctx.Param("token")

	client := http.Client{}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s/guacamole/api/tokens/%s", c.Url, token), nil)
	if err != nil {
		return ctx.JSON(http.StatusNoContent, model.SuccessResponse{
			StatusCode: http.StatusNoContent,
			Message:    "error delete token",
			Data:       false,
		})
	}
	response, responseError := client.Do(req)
	if responseError != nil {
		return ctx.JSON(http.StatusNoContent, model.SuccessResponse{
			StatusCode: http.StatusNoContent,
			Message:    "error response delete token",
		})
	}
	defer response.Body.Close()

	return ctx.JSON(http.StatusOK, model.SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    "success delete token",
		Data:       true,
	})
}
