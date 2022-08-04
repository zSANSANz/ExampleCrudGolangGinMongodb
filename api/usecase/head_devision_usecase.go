package crud

import (
	"context"

	"chatnews-api/api/model"
)

type HeadDevisionUseCaseI interface {
	GetDataUC(ctx context.Context) (resp model.GetDataHeadDevisionResponse, err error)
	InsertDataUC(ctx context.Context, req model.DataHeadDevisionRequest) (resp bool, err error)
	UpdateDataUC(ctx context.Context, req model.DataHeadDevisionRequest) (resp bool, err error)
	DeleteDataUC(ctx context.Context, req model.DataHeadDevisionRequest) (resp bool, err error)
}
