package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORSConfig(e *echo.Echo) {
	config := middleware.CORSConfig{
		// AllowHeaders:     []string{"*"},
		// AllowCredentials: true,
		// AllowOrigins:     []string{"*"},
		// AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}
	e.Use(middleware.CORSWithConfig(config))
}
