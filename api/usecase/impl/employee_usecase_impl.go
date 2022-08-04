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

type EmployeeUseCase struct {
	config       *model.EnvConfig
	employeeRepo repository.EmployeeRepositoryI
}

func NewEmployeeUseCase(config *model.EnvConfig, employeeRepo repository.EmployeeRepositoryI) usecase.EmployeeUseCaseI {
	return &EmployeeUseCase{
		config:       config,
		employeeRepo: employeeRepo,
	}
}

func (cuc *EmployeeUseCase) GetDataUC(ctx context.Context) (resp model.GetDataEmployeeResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := cuc.employeeRepo.GetAllDataEmployee(ctx)
	if err != nil {
		log.Println("failed to show data employee with default log")
		return list, err
	}

	return list, err
}

func (cuc *EmployeeUseCase) InsertDataUC(ctx context.Context, req model.DataEmployeeRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	if req.Name == "" {
		err = errors.New("failed to add data employee ")
		logging.Info(err)
		return false, err
	}

	//insert data
	err = cuc.employeeRepo.InsertDataEmployee(ctx, req)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (cuc *EmployeeUseCase) UpdateDataUC(ctx context.Context, req model.DataEmployeeRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	err = cuc.employeeRepo.UpdateDataEmployee(ctx, req)
	if err != nil {
		return false, err
	}
	zaplogger, _ := zap.NewProduction()
	defer zaplogger.Sync()
	zaplogger.Info("success to update product",
		zap.String("employee ID", req.ID.Hex()),
	)

	return true, nil
}

func (cuc *EmployeeUseCase) DeleteDataUC(ctx context.Context, req model.DataEmployeeRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	err = cuc.employeeRepo.DeleteDataEmployee(ctx, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
