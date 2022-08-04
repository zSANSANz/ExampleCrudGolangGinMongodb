package crud

import (
	"context"

	"chatnews-api/api/model"

	"github.com/labstack/echo/v4"
)

type PerformanceRepositoryI interface {
	GetAllDataPerformance(ctx context.Context) (crudResp model.GetDataPerformanceResponse, err error)
	GetAllDataPerformanceLikeDevision(ctx context.Context) (crudResp model.GetDataPerformanceResponse, err error)
	GetAllDataPerformanceLikeUser(ctx context.Context) (crudResp model.GetDataPerformanceResponse, err error)
	GetAllDataPerformanceOrderByUpdatedAt(ctx context.Context) (crudResp model.GetDataPerformanceResponse, err error)
	InsertDataPerformance(c echo.Context, ctx context.Context, req model.DataPerformanceRequest) error
	UpdateDataPerformance(c echo.Context, ctx context.Context, req model.DataPerformanceRequest) error
	UpdateStatusPerformance(c echo.Context, ctx context.Context, req model.DataPerformanceRequest) error
	DeleteDataPerformance(c echo.Context, ctx context.Context, req model.DataPerformanceRequest) error
}
