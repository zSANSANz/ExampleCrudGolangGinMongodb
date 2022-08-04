package usecase

import (
	"context"
	"log"

	"chatnews-api/api/model"
	repository "chatnews-api/api/repository"
	usecase "chatnews-api/api/usecase"
	"chatnews-api/lib/logging"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type HeadDevisionUseCase struct {
	config           *model.EnvConfig
	headDevisionRepo repository.HeadDevisionRepositoryI
}

func NewHeadDevisionUseCase(config *model.EnvConfig, headDevisionRepo repository.HeadDevisionRepositoryI) usecase.HeadDevisionUseCaseI {
	return &HeadDevisionUseCase{
		config:           config,
		headDevisionRepo: headDevisionRepo,
	}
}

func (cuc *HeadDevisionUseCase) GetDataUC(ctx context.Context) (resp model.GetDataHeadDevisionResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := cuc.headDevisionRepo.GetAllDataHeadDevision(ctx)
	if err != nil {
		log.Println("failed to show data headDevision with default log")
		return list, err
	}

	return list, err
}

func (cuc *HeadDevisionUseCase) InsertDataUC(ctx context.Context, req model.DataHeadDevisionRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	if req.Name == "" {
		err = errors.New("failed to add data headDevision ")
		logging.Info(err)
		return false, err
	}

	//insert data
	err = cuc.headDevisionRepo.InsertDataHeadDevision(ctx, req)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (cuc *HeadDevisionUseCase) UpdateDataUC(ctx context.Context, req model.DataHeadDevisionRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	err = cuc.headDevisionRepo.UpdateDataHeadDevision(ctx, req)
	if err != nil {
		return false, err
	}
	zaplogger, _ := zap.NewProduction()
	defer zaplogger.Sync()
	zaplogger.Info("success to update product",
		zap.String("headDevision ID", req.ID.Hex()),
	)

	return true, nil
}

func (cuc *HeadDevisionUseCase) DeleteDataUC(ctx context.Context, req model.DataHeadDevisionRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	err = cuc.headDevisionRepo.DeleteDataHeadDevision(ctx, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
