package crud

import (
	"context"

	"chatnews-api/api/model"
)

type EmployeeRepositoryI interface {
	GetAllDataEmployee(ctx context.Context) (crudResp model.GetDataEmployeeResponse, err error)
	InsertDataEmployee(ctx context.Context, req model.DataEmployeeRequest) error
	UpdateDataEmployee(ctx context.Context, req model.DataEmployeeRequest) error
	DeleteDataEmployee(ctx context.Context, req model.DataEmployeeRequest) error
}
