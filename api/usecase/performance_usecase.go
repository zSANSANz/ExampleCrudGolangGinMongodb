package crud

import (
	"context"

	"chatnews-api/api/model"

	"github.com/labstack/echo/v4"
)

type PerformanceUseCaseI interface {
	GetDataUC(ctx context.Context) (resp model.GetDataPerformanceResponse, err error)
	GetDataUCUser(ctx context.Context) (resp model.GetDataPerformanceResponse, err error)
	GetDataUCDevision(ctx context.Context) (resp model.GetDataPerformanceResponse, err error)
	GetDataUCOrderByUpdatedAt(ctx context.Context) (resp model.GetDataPerformanceResponse, err error)
	InsertDataUC(c echo.Context, ctx context.Context, req model.DataPerformanceRequest) (resp bool, err error)
	UpdateDataUC(c echo.Context, ctx context.Context, req model.DataPerformanceRequest) (resp bool, err error)
	UpdateStatusUC(c echo.Context, ctx context.Context, req model.DataPerformanceRequest) (resp bool, err error)
	DeleteDataUC(c echo.Context, ctx context.Context, req model.DataPerformanceRequest) (resp bool, err error)
}
