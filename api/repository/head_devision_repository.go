package crud

import (
	"context"

	"chatnews-api/api/model"
)

type HeadDevisionRepositoryI interface {
	GetAllDataHeadDevision(ctx context.Context) (crudResp model.GetDataHeadDevisionResponse, err error)
	InsertDataHeadDevision(ctx context.Context, req model.DataHeadDevisionRequest) error
	UpdateDataHeadDevision(ctx context.Context, req model.DataHeadDevisionRequest) error
	DeleteDataHeadDevision(ctx context.Context, req model.DataHeadDevisionRequest) error
}
