package repository

import (
	"chatnews-api/middleware/exception"
	"chatnews-api/middleware/model"

	paginate "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PerformanceRepository interface {
	GetAllPerformance(page int64, limit int64) (*model.PagedPerformance, error)
	GetPerformanceHod(id string) ([]model.Performance, error)
	GetPerformanceSupervision(id string) ([]model.Performance, error)
	GetPerformanceUser(id string) ([]model.Performance, error)
	GetPerformanceUserId(id string) ([]model.Performance, error)
	GetPerformanceId(id string) ([]model.Performance, error)
}

type performanceRepositoryImpl struct {
	Connection *mongo.Database
}

func NewPerformanceRepository(Connection *mongo.Database) PerformanceRepository {
	return &performanceRepositoryImpl{Connection: Connection}
}

func (PerformanceRepository *performanceRepositoryImpl) GetAllPerformance(page int64, limit int64) (*model.PagedPerformance, error) {
	var performances []model.Performance

	filter := bson.D{}

	findOptions := options.Find()

	findOptions.SetSort(bson.D{{"created_at", -1}})

	collection := PerformanceRepository.Connection.Collection("performances")

	projection := bson.D{
		{"id", 1},
		{"employee", 1},
		{"division", 1},
		{"status", 1},
		{"task_productivity", 1},
		{"task_personal", 1},
		{"kpi_status", 1},
		{"noted", 1},
		{"total_score", 1},
		{"created_by", 1},
		{"created_at", 1},
		{"update_by", 1},
		{"update_at", 1},
		{"deleted_by", 1},
		{"deleted_at", 1},
	}

	paginatedData, err := paginate.New(collection).Context(cntx).Limit(limit).Page(page).Select(projection).Filter(filter).Decode(&performances).Find()
	if err != nil {
		return nil, err
	}

	return &model.PagedPerformance{
		Data:     performances,
		PageInfo: paginatedData.Pagination,
	}, nil
}

func (performanceRepository *performanceRepositoryImpl) GetPerformanceHod(id string) ([]model.Performance, error) {
	var existingPerformances []model.Performance
	filter := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "id_employee",
				"foreignField": "_id",
				"as":           "employee",
			},
		},
		bson.M{
			"$match": bson.M{
				"deleted": bson.M{
					"$ne": true,
				},
				"division": bson.M{"$regex": id},
			},
		},
		bson.M{
			"$sort": bson.M{
				"created_at": -1,
			},
		},
	}

	cursor, err := performanceRepository.Connection.Collection("performances").Aggregate(cntx, filter)
	if err != nil {
		return nil, exception.ResourceNotFoundException("Performance", "division", id)
	}

	if err = cursor.All(cntx, &existingPerformances); err != nil {
		return nil, exception.ResourceNotFoundException("Performance", "division", id)
	}

	return existingPerformances, nil
}

func (performanceRepository *performanceRepositoryImpl) GetPerformanceSupervision(id string) ([]model.Performance, error) {
	var existingPerformances []model.Performance
	filter := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "id_employee",
				"foreignField": "_id",
				"as":           "employee",
			},
		},
		bson.M{
			"$match": bson.M{
				"deleted": bson.M{
					"$ne": true,
				},
				"employee.superior": bson.M{"$regex": id},
			},
		},
		bson.M{
			"$sort": bson.M{
				"created_at": -1,
			},
		},
	}

	cursor, err := performanceRepository.Connection.Collection("performances").Aggregate(cntx, filter)
	if err != nil {
		return nil, exception.ResourceNotFoundException("Performance", "superior", id)
	}

	if err = cursor.All(cntx, &existingPerformances); err != nil {
		return nil, exception.ResourceNotFoundException("Performance", "superior", id)
	}

	return existingPerformances, nil
}

func (performanceRepository *performanceRepositoryImpl) GetPerformanceUser(id string) ([]model.Performance, error) {
	var existingPerformances []model.Performance
	filter := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "id_employee",
				"foreignField": "_id",
				"as":           "employee",
			},
		},
		bson.M{
			"$match": bson.M{
				"deleted": bson.M{
					"$ne": true,
				},
				"employee.firstName": bson.M{"$regex": id},
			},
		},
		bson.M{
			"$sort": bson.M{
				"created_at": -1,
			},
		},
	}

	cursor, err := performanceRepository.Connection.Collection("performances").Aggregate(cntx, filter)
	if err != nil {
		return nil, exception.ResourceNotFoundException("Performance", "id_employee", id)
	}

	if err = cursor.All(cntx, &existingPerformances); err != nil {
		return nil, exception.ResourceNotFoundException("Performance", "id_employee", id)
	}

	return existingPerformances, nil
}

func (performanceRepository *performanceRepositoryImpl) GetPerformanceUserId(id string) ([]model.Performance, error) {
	var existingPerformances []model.Performance
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "id_employee",
				"foreignField": "_id",
				"as":           "employee",
			},
		},
		bson.M{
			"$match": bson.M{
				"deleted": bson.M{
					"$ne": true,
				},
				"employee._id": objectId,
			},
		},
		bson.M{
			"$sort": bson.M{
				"created_at": -1,
			},
		},
	}

	cursor, err := performanceRepository.Connection.Collection("performances").Aggregate(cntx, filter)
	if err != nil {
		return nil, exception.ResourceNotFoundException("Performance", "id_employee", id)
	}

	if err = cursor.All(cntx, &existingPerformances); err != nil {
		return nil, exception.ResourceNotFoundException("Performance", "id_employee", id)
	}

	return existingPerformances, nil
}

func (performanceRepository *performanceRepositoryImpl) GetPerformanceId(id string) ([]model.Performance, error) {
	var existingPerformance []model.Performance
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "id_employee",
				"foreignField": "_id",
				"as":           "employee",
			},
		},
		bson.M{
			"$match": bson.M{
				"deleted": bson.M{
					"$ne": true,
				},
				"_id": objectId,
			},
		},
		bson.M{
			"$sort": bson.M{
				"created_at": -1,
			},
		},
	}
	cursor, err := performanceRepository.Connection.Collection("performances").Aggregate(cntx, filter)

	if err != nil {
		return nil, exception.ResourceNotFoundException("Performance", "id", id)
	}

	if err = cursor.All(cntx, &existingPerformance); err != nil {
		return nil, exception.ResourceNotFoundException("Performance", "id_performance", id)
	}

	return existingPerformance, nil
}
