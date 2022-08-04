package usecase

import (
	"context"
	"log"

	"chatnews-api/api/model"
	repository "chatnews-api/api/repository"
	usecase "chatnews-api/api/usecase"
	"chatnews-api/lib/logging"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type PerformanceUseCase struct {
	config          *model.EnvConfig
	performanceRepo repository.PerformanceRepositoryI
}

func NewPerformanceUseCase(config *model.EnvConfig, performanceRepo repository.PerformanceRepositoryI) usecase.PerformanceUseCaseI {
	return &PerformanceUseCase{
		config:          config,
		performanceRepo: performanceRepo,
	}
}

func (cuc *PerformanceUseCase) GetDataUC(ctx context.Context) (resp model.GetDataPerformanceResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := cuc.performanceRepo.GetAllDataPerformance(ctx)
	if err != nil {
		log.Println("failed to show data performance with default log")
		return list, err
	}

	return list, err
}

func (cuc *PerformanceUseCase) GetDataUCUser(ctx context.Context) (resp model.GetDataPerformanceResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := cuc.performanceRepo.GetAllDataPerformanceLikeUser(ctx)
	if err != nil {
		log.Println("failed to show data performance with default log")
		return list, err
	}

	return list, err
}

func (cuc *PerformanceUseCase) GetDataUCDevision(ctx context.Context) (resp model.GetDataPerformanceResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := cuc.performanceRepo.GetAllDataPerformanceLikeDevision(ctx)
	if err != nil {
		log.Println("failed to show data performance with default log")
		return list, err
	}

	return list, err
}

func (cuc *PerformanceUseCase) GetDataUCOrderByUpdatedAt(ctx context.Context) (resp model.GetDataPerformanceResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := cuc.performanceRepo.GetAllDataPerformanceOrderByUpdatedAt(ctx)
	if err != nil {
		log.Println("failed to show data performance with default log")
		return list, err
	}

	return list, err
}

func (cuc *PerformanceUseCase) InsertDataUC(c echo.Context, ctx context.Context, req model.DataPerformanceRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	if req.Status == "" {
		err = errors.New("failed to add data performance ")
		logging.Info(err)
		return false, err
	}

	//insert data
	err = cuc.performanceRepo.InsertDataPerformance(c, ctx, req)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (cuc *PerformanceUseCase) UpdateDataUC(c echo.Context, ctx context.Context, req model.DataPerformanceRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	err = cuc.performanceRepo.UpdateDataPerformance(c, ctx, req)
	if err != nil {
		return false, err
	}
	zaplogger, _ := zap.NewProduction()
	defer zaplogger.Sync()
	zaplogger.Info("success to update product",
		zap.String("performance ID", req.ID.Hex()),
	)

	return true, nil
}

func (cuc *PerformanceUseCase) UpdateStatusUC(c echo.Context, ctx context.Context, req model.DataPerformanceRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	err = cuc.performanceRepo.UpdateStatusPerformance(c, ctx, req)
	if err != nil {
		return false, err
	}
	zaplogger, _ := zap.NewProduction()
	defer zaplogger.Sync()
	zaplogger.Info("success to update product",
		zap.String("performance ID", req.ID.Hex()),
	)

	return true, nil
}

func (cuc *PerformanceUseCase) DeleteDataUC(c echo.Context, ctx context.Context, req model.DataPerformanceRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	err = cuc.performanceRepo.DeleteDataPerformance(c, ctx, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
