package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"chatnews-api/api/model"
	repository "chatnews-api/api/repository"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepository struct {
	mongoDB *mongo.Database
}

func NewEmployeeRepository(mongo *mongo.Database) repository.EmployeeRepositoryI {
	return &EmployeeRepository{
		mongoDB: mongo,
	}
}

func (cr EmployeeRepository) GetAllDataEmployee(ctx context.Context) (crudResp model.GetDataEmployeeResponse, err error) {

	query, err := cr.mongoDB.Collection("employees").Find(ctx, bson.D{})
	if err != nil {
		log.Println("error", err)
		return model.GetDataEmployeeResponse{}, err
	}
	defer query.Close(ctx)

	listDataEmployee := make([]model.DataEmployee, 0)
	for query.Next(ctx) {
		var row model.DataEmployee
		err := query.Decode(&row)
		if err != nil {
			log.Println("error")
		}
		listDataEmployee = append(listDataEmployee, row)
	}

	crudResp = model.GetDataEmployeeResponse{Data: listDataEmployee}

	return crudResp, err
}

func (cr EmployeeRepository) InsertDataEmployee(ctx context.Context, req model.DataEmployeeRequest) (err error) {
	// var cthAsia, _ = time.LoadLocation("Asia/Bangkok")

	dataReq := bson.M{
		"name":       req.Name,
		"created_at": time.Now(),
	}

	query, err := cr.mongoDB.Collection("employees").InsertOne(ctx, dataReq)
	if err != nil {
		log.Println("error")
	}

	if oid, ok := query.InsertedID.(primitive.ObjectID); ok {
		employeeID := oid.Hex()
		dataUpdateEmployeeID := bson.M{"_id": oid}
		dataObjectID := bson.M{"$set": bson.M{"_id": employeeID}}
		_, err := cr.mongoDB.Collection("employees").UpdateOne(ctx, dataUpdateEmployeeID, dataObjectID)
		if err != nil {
			log.Println("error")
		}
	} else {
		err = errors.New(fmt.Sprint("can't get inserted ID ", err))
		log.Println("error")
	}

	return err
}

func (cr EmployeeRepository) UpdateDataEmployee(ctx context.Context, req model.DataEmployeeRequest) (err error) {

	dataUpdateEmployeeID := bson.M{"_id": req.ID}
	dataObjectID := bson.M{"$set": bson.M{
		"name":       req.Name,
		"updated_at": time.Now(),
	}}
	_, err = cr.mongoDB.Collection("employees").UpdateOne(ctx, dataUpdateEmployeeID, dataObjectID)
	if err != nil {
		log.Println("error")
	}

	return err
}

func (cr EmployeeRepository) DeleteDataEmployee(ctx context.Context, req model.DataEmployeeRequest) (err error) {

	dataUpdateEmployeeID := bson.M{"_id": req.ID}

	_, err = cr.mongoDB.Collection("employees").DeleteOne(ctx, dataUpdateEmployeeID)
	if err != nil {
		log.Println("error")
	}

	return err
}
