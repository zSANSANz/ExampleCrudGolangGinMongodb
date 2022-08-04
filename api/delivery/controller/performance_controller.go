package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"chatnews-api/api/model"
	api "chatnews-api/api/usecase"

	"github.com/labstack/echo/v4"
)

type PerformanceController struct {
	e       *echo.Echo
	usecase api.PerformanceUseCaseI
}

func NewPerformanceController(e *echo.Echo, usecase api.PerformanceUseCaseI) *PerformanceController {
	return &PerformanceController{
		e:       e,
		usecase: usecase,
	}
}

func (cc *PerformanceController) GetData(ec echo.Context) error {

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

func (cc *PerformanceController) GetDataUCUser(ec echo.Context) error {

	data, err := cc.usecase.GetDataUCUser(context.Background())
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

func (cc *PerformanceController) GetDataUCDevision(ec echo.Context) error {

	data, err := cc.usecase.GetDataUCDevision(context.Background())
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

func (cc *PerformanceController) GetDataUCOrderByUpdatedAt(ec echo.Context) error {

	data, err := cc.usecase.GetDataUCOrderByUpdatedAt(context.Background())
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

func (cc *PerformanceController) InsertData(ec echo.Context) error {

	var req model.DataPerformanceRequest

	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       err,
		})
	}

	data, err := cc.usecase.InsertDataUC(ec, context.Background(), req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       data,
		})
	}

	return nil
}

func (cc *PerformanceController) UpdateData(ec echo.Context) error {

	var req model.DataPerformanceRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       "",
		})
	}
	data, err := cc.usecase.UpdateDataUC(ec, context.Background(), req)
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

func (cc *PerformanceController) UpdateStatus(ec echo.Context) error {

	var req model.DataPerformanceRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       "",
		})
	}
	data, err := cc.usecase.UpdateStatusUC(ec, context.Background(), req)
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

func (cc *PerformanceController) DeleteData(ec echo.Context) error {

	var req model.DataPerformanceRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       "",
		})
	}
	data, err := cc.usecase.DeleteDataUC(ec, context.Background(), req)
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
