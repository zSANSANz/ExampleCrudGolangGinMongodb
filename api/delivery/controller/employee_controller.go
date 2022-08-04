package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"chatnews-api/api/model"
	api "chatnews-api/api/usecase"

	"github.com/labstack/echo/v4"
)

type EmployeeController struct {
	e       *echo.Echo
	usecase api.EmployeeUseCaseI
}

func NewEmployeeController(e *echo.Echo, usecase api.EmployeeUseCaseI) *EmployeeController {
	return &EmployeeController{
		e:       e,
		usecase: usecase,
	}
}

func (cc *EmployeeController) GetData(ec echo.Context) error {

	data, err := cc.usecase.GetDataUC(context.Background())
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       data,
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success",
		"data":    data,
	})
}

func (cc *EmployeeController) InsertData(ec echo.Context) error {

	var req model.DataEmployeeRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       "",
		})
	}
	data, err := cc.usecase.InsertDataUC(context.Background(), req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       data,
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success",
		"data":    data,
	})
}

func (cc *EmployeeController) UpdateData(ec echo.Context) error {

	var req model.DataEmployeeRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       "",
		})
	}
	data, err := cc.usecase.UpdateDataUC(context.Background(), req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       data,
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success",
		"data":    data,
	})
}

func (cc *EmployeeController) DeleteData(ec echo.Context) error {

	var req model.DataEmployeeRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       "",
		})
	}
	data, err := cc.usecase.DeleteDataUC(context.Background(), req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       data,
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success",
		"data":    data,
	})
}
