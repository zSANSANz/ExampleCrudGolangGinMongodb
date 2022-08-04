package crud

import (
	"context"

	"chatnews-api/api/model"
)

type EmployeeUseCaseI interface {
	GetDataUC(ctx context.Context) (resp model.GetDataEmployeeResponse, err error)
	InsertDataUC(ctx context.Context, req model.DataEmployeeRequest) (resp bool, err error)
	UpdateDataUC(ctx context.Context, req model.DataEmployeeRequest) (resp bool, err error)
	DeleteDataUC(ctx context.Context, req model.DataEmployeeRequest) (resp bool, err error)
}
