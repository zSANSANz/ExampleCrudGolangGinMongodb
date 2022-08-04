package controller

import (
	"chatnews-api/middleware/repository"
	"chatnews-api/middleware/security"
	"chatnews-api/middleware/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PerformanceController struct {
	performanceRepository repository.PerformanceRepository
	authValidator         *security.AuthValidator
}

func NewPerformanceController(performanceRepository repository.PerformanceRepository, authValidator *security.AuthValidator) *PerformanceController {
	return &PerformanceController{performanceRepository: performanceRepository, authValidator: authValidator}
}

func (performanceController *PerformanceController) GetAllPerformance(c echo.Context) error {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)

	pagedPerformance, _ := performanceController.performanceRepository.GetAllPerformance(page, limit)
	return util.Negotiate(c, http.StatusOK, pagedPerformance)
}

func (performanceController *PerformanceController) GetPerformanceHod(c echo.Context) error {
	id := c.Param("id")

	performance, err := performanceController.performanceRepository.GetPerformanceHod(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success",
		"data":    performance,
	})
}

func (performanceController *PerformanceController) GetPerformanceSupervision(c echo.Context) error {
	id := c.Param("id")

	performance, err := performanceController.performanceRepository.GetPerformanceSupervision(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success",
		"data":    performance,
	})
}

func (performanceController *PerformanceController) GetPerformanceUser(c echo.Context) error {
	id := c.Param("id")

	performance, err := performanceController.performanceRepository.GetPerformanceUser(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success",
		"data":    performance,
	})
}

func (performanceController *PerformanceController) GetPerformanceUserId(c echo.Context) error {
	id := c.Param("id")

	performance, err := performanceController.performanceRepository.GetPerformanceUserId(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success",
		"data":    performance,
	})
}

func (performanceController *PerformanceController) GetPerformanceId(c echo.Context) error {
	id := c.Param("id")

	performance, err := performanceController.performanceRepository.GetPerformanceId(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success",
		"data":    performance,
	})
}
