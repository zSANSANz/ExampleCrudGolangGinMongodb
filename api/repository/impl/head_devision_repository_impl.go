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

type HeadDevisionRepository struct {
	mongoDB *mongo.Database
}

func NewHeadDevisionRepository(mongo *mongo.Database) repository.HeadDevisionRepositoryI {
	return &HeadDevisionRepository{
		mongoDB: mongo,
	}
}

func (cr HeadDevisionRepository) GetAllDataHeadDevision(ctx context.Context) (crudResp model.GetDataHeadDevisionResponse, err error) {

	query, err := cr.mongoDB.Collection("head_devisions").Find(ctx, bson.D{})
	if err != nil {
		log.Println("error", err)
		return model.GetDataHeadDevisionResponse{}, err
	}
	defer query.Close(ctx)

	listDataHeadDevision := make([]model.DataHeadDevision, 0)
	for query.Next(ctx) {
		var row model.DataHeadDevision
		err := query.Decode(&row)
		if err != nil {
			log.Println("error")
		}
		listDataHeadDevision = append(listDataHeadDevision, row)
	}

	crudResp = model.GetDataHeadDevisionResponse{Data: listDataHeadDevision}

	return crudResp, err
}

func (cr HeadDevisionRepository) InsertDataHeadDevision(ctx context.Context, req model.DataHeadDevisionRequest) (err error) {
	// var cthAsia, _ = time.LoadLocation("Asia/Bangkok")

	dataReq := bson.M{
		"name":       req.Name,
		"created_at": time.Now(),
	}

	query, err := cr.mongoDB.Collection("head_devisions").InsertOne(ctx, dataReq)
	if err != nil {
		log.Println("error")
	}

	if oid, ok := query.InsertedID.(primitive.ObjectID); ok {
		headDevisionID := oid.Hex()
		dataUpdateHeadDevisionID := bson.M{"_id": oid}
		dataObjectID := bson.M{"$set": bson.M{"_id": headDevisionID}}
		_, err := cr.mongoDB.Collection("head_devisions").UpdateOne(ctx, dataUpdateHeadDevisionID, dataObjectID)
		if err != nil {
			log.Println("error")
		}
	} else {
		err = errors.New(fmt.Sprint("can't get inserted ID ", err))
		log.Println("error")
	}

	return err
}

func (cr HeadDevisionRepository) UpdateDataHeadDevision(ctx context.Context, req model.DataHeadDevisionRequest) (err error) {

	dataUpdateHeadDevisionID := bson.M{"_id": req.ID}
	dataObjectID := bson.M{"$set": bson.M{
		"name":       req.Name,
		"updated_at": time.Now(),
	}}
	_, err = cr.mongoDB.Collection("head_devisions").UpdateOne(ctx, dataUpdateHeadDevisionID, dataObjectID)
	if err != nil {
		log.Println("error")
	}

	return err
}

func (cr HeadDevisionRepository) DeleteDataHeadDevision(ctx context.Context, req model.DataHeadDevisionRequest) (err error) {

	dataUpdateHeadDevisionID := bson.M{"_id": req.ID}

	_, err = cr.mongoDB.Collection("head_devisions").DeleteOne(ctx, dataUpdateHeadDevisionID)
	if err != nil {
		log.Println("error")
	}

	return err
}
