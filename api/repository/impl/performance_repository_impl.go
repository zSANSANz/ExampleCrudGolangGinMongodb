package repository

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"
	"time"

	"chatnews-api/api/model"
	repository "chatnews-api/api/repository"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PerformanceRepository struct {
	mongoDB *mongo.Database
}

func NewPerformanceRepository(mongo *mongo.Database) repository.PerformanceRepositoryI {
	return &PerformanceRepository{
		mongoDB: mongo,
	}
}

func (cr PerformanceRepository) GetAllDataPerformance(ctx context.Context) (crudResp model.GetDataPerformanceResponse, err error) {

	lookupStage := bson.A{
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
			},
		},
		bson.M{
			"$sort": bson.M{
				"created_at": -1,
			},
		},
	}

	query, err := cr.mongoDB.Collection("performances").Aggregate(ctx, lookupStage)

	if err != nil {
		log.Println("error", err)
		return model.GetDataPerformanceResponse{}, err
	}
	defer query.Close(ctx)

	listDataPerformance := make([]model.DataPerformance, 0)
	for query.Next(ctx) {
		var row model.DataPerformance
		err := query.Decode(&row)
		if err != nil {
			log.Println("error")
		}
		listDataPerformance = append(listDataPerformance, row)
	}

	crudResp = model.GetDataPerformanceResponse{Data: listDataPerformance}

	return crudResp, err
}

func (cr PerformanceRepository) GetAllDataPerformanceLikeDevision(ctx context.Context) (crudResp model.GetDataPerformanceResponse, err error) {

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"created_at", -1}})

	query, err := cr.mongoDB.Collection("performances").Find(ctx, bson.M{"division": bson.M{"$regex": `sy.*`}}, findOptions)
	if err != nil {
		log.Println("error", err)
		return model.GetDataPerformanceResponse{}, err
	}
	defer query.Close(ctx)

	listDataPerformance := make([]model.DataPerformance, 0)
	for query.Next(ctx) {
		var row model.DataPerformance
		err := query.Decode(&row)
		if err != nil {
			log.Println("error")
		}
		listDataPerformance = append(listDataPerformance, row)
	}

	crudResp = model.GetDataPerformanceResponse{Data: listDataPerformance}

	return crudResp, err
}

func (cr PerformanceRepository) GetAllDataPerformanceLikeUser(ctx context.Context) (crudResp model.GetDataPerformanceResponse, err error) {

	findOptions := options.Find()
	// Sort by `price` field descending
	findOptions.SetSort(bson.D{{"created_at", -1}})

	query, err := cr.mongoDB.Collection("performances").Find(ctx, bson.M{"employee.firstName": bson.M{"$regex": `sa.*`}}, findOptions)
	if err != nil {
		log.Println("error", err)
		return model.GetDataPerformanceResponse{}, err
	}
	defer query.Close(ctx)

	listDataPerformance := make([]model.DataPerformance, 0)
	for query.Next(ctx) {
		var row model.DataPerformance
		err := query.Decode(&row)
		if err != nil {
			log.Println("error")
		}
		listDataPerformance = append(listDataPerformance, row)
	}

	crudResp = model.GetDataPerformanceResponse{Data: listDataPerformance}

	return crudResp, err
}

func (cr PerformanceRepository) GetAllDataPerformanceOrderByUpdatedAt(ctx context.Context) (crudResp model.GetDataPerformanceResponse, err error) {

	// findOptions := options.Find()
	// findOptions.SetSort(bson.D{{"update_at", -1}})

	lookupStage := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "id_employee",
				"foreignField": "_id",
				"as":           "employee",
			},
		},
	}

	query, err := cr.mongoDB.Collection("performances").Aggregate(ctx, lookupStage)

	if err != nil {
		log.Println("error", err)
		return model.GetDataPerformanceResponse{}, err
	}
	defer query.Close(ctx)

	listDataPerformance := make([]model.DataPerformance, 0)
	for query.Next(ctx) {
		var row model.DataPerformance
		err := query.Decode(&row)
		if err != nil {
			log.Println("error")
		}
		listDataPerformance = append(listDataPerformance, row)
	}

	crudResp = model.GetDataPerformanceResponse{Data: listDataPerformance}

	return crudResp, err
}

func (cr PerformanceRepository) InsertDataPerformance(c echo.Context, ctx context.Context, req model.DataPerformanceRequest) (err error) {

	tokenString := ""

	// get jwt bearer token from context
	for key, values := range c.Request().Header {
		for _, value := range values {
			if key == "Authorization" {
				tokenString = strings.Replace(value, "Bearer ", "", -1)

			}
		}
	}

	// decode token with secret code bermaslaah
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("bermaslaah"), nil
	})

	if !token.Valid {
		fmt.Println("invalid jwt token!!!")
	}

	var existingUsers model.DataUser
	objectId, _ := primitive.ObjectIDFromHex(claims["id"].(string))
	filter := bson.M{"_id": objectId}

	errUser := cr.mongoDB.Collection("users").FindOne(ctx, filter).Decode(&existingUsers)

	if errUser != nil {
		log.Println("error")
	}

	var existingEmployees model.DataUserPerformance
	// objectIdEmp, _ := primitive.ObjectIDFromHex(req.IdEmployee)
	objectIdEmp := req.IdEmployee
	filterEmp := bson.M{"_id": objectIdEmp}

	errEmployees := cr.mongoDB.Collection("users").FindOne(ctx, filterEmp).Decode(&existingEmployees)

	if errEmployees != nil {
		log.Println("error")
	}
	fmt.Println(existingEmployees)

	dataReq := bson.M{
		"id_employee":       req.IdEmployee,
		"status":            req.Status,
		"task_productivity": req.TaskProductivity,
		"task_personal":     req.TaskPersonal,
		"kpi_status":        req.KpiStatus,
		"noted":             req.Noted,
		"total_score":       req.TotalScore,
		"score_category":    req.ScoreCategory,
		"created_at":        time.Now(),
		"created_by":        existingUsers.IdEmp,
		"created_by_id":     claims["id"].(string),
	}

	if !existingUsers.Head {
		// fmt.Println("kamu bukanlah head")
		fmt.Println(existingUsers)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "Unauthorized: Access is denied due to invalid credentials",
			"error_code": "401",
			"data":       existingUsers,
		})
	} else {
		// fmt.Println("kamu adalah head")
		query, err := cr.mongoDB.Collection("performances").InsertOne(ctx, dataReq)
		if err != nil {
			log.Println("error")
		}

		if oid, ok := query.InsertedID.(primitive.ObjectID); ok {
			performanceID := oid.Hex()
			dataUpdatePerformanceID := bson.M{"_id": oid}
			dataObjectID := bson.M{"$set": bson.M{"_id": performanceID}}
			_, err := cr.mongoDB.Collection("performances").UpdateOne(ctx, dataUpdatePerformanceID, dataObjectID)
			if err != nil {
				log.Println("error")
			}
		} else {
			err = errors.New(fmt.Sprint("can't get inserted ID ", err))
			log.Println("error")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "true",
			"message": "Success",
			"data":    existingUsers,
		})
	}

}

func (cr PerformanceRepository) UpdateDataPerformance(c echo.Context, ctx context.Context, req model.DataPerformanceRequest) (err error) {

	tokenString := ""

	// get jwt bearer token from context
	for key, values := range c.Request().Header {
		for _, value := range values {
			if key == "Authorization" {
				tokenString = strings.Replace(value, "Bearer ", "", -1)

			}
		}
	}

	// decode token with secret code bermaslaah
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("bermaslaah"), nil
	})

	if !token.Valid {
		fmt.Println("invalid jwt token!!!")
	}

	dataUpdatePerformanceID := bson.M{"_id": req.ID}

	var reqPerformance model.DataPerformance
	filter := bson.M{"_id": req.ID}
	errPerformance := cr.mongoDB.Collection("performances").FindOne(ctx, filter).Decode(&reqPerformance)

	if errPerformance != nil {
		log.Println("error")
	}

	RevisionCount := reqPerformance.RevisionCount
	if req.Status == "review" {
		RevisionCount += 1
	}

	var reqUser model.DataUser
	objectId, _ := primitive.ObjectIDFromHex(claims["id"].(string))
	filterUser := bson.M{"_id": objectId}
	errUser := cr.mongoDB.Collection("users").FindOne(ctx, filterUser).Decode(&reqUser)

	if errUser != nil {
		log.Println("error")
	}

	var existingEmployees model.DataUser
	objectIdEmp := req.IdEmployee
	filterEmp := bson.M{"_id": objectIdEmp}

	errEmployees := cr.mongoDB.Collection("users").FindOne(ctx, filterEmp).Decode(&existingEmployees)

	if errEmployees != nil {
		log.Println("error")
	}
	fmt.Println(existingEmployees)

	dataObjectID := bson.M{"$set": bson.M{
		"id_employee":       req.IdEmployee,
		"status":            req.Status,
		"task_productivity": req.TaskProductivity,
		"task_personal":     req.TaskPersonal,
		"kpi_status":        req.KpiStatus,
		"noted":             req.Noted,
		"total_score":       req.TotalScore,
		"score_category":    req.ScoreCategory,
		"revision_count":    RevisionCount,
		"update_at":         time.Now(),
		"update_by":         reqUser.IdEmp,
		"update_by_id":      claims["id"].(string),
	}}
	_, err = cr.mongoDB.Collection("performances").UpdateOne(ctx, dataUpdatePerformanceID, dataObjectID)
	if err != nil {
		log.Println("error")
	}

	fmt.Println(existingEmployees, "error", "iha.pariha@indochat.co.id")

	// to := []string{existingEmployees.Email}
	// cc := []string{"iha.pariha@indochat.co.id"}
	// subject := "KPI Performance"
	// message := fmt.Sprintf("Hello %s Please Check your KPI performance on %s division, the employee based on named %s, the status has been changed to %s", existingEmployees.Superior, existingEmployees.Division, existingEmployees.FirstName, req.Status)

	// errMail := sendMail(to, cc, subject, message)
	// if errMail != nil {
	// 	log.Fatal(errMail.Error())
	// }

	log.Println("Mail sent!")

	return err
}

func (cr PerformanceRepository) UpdateStatusPerformance(c echo.Context, ctx context.Context, req model.DataPerformanceRequest) (err error) {

	tokenString := ""

	// get jwt bearer token from context
	for key, values := range c.Request().Header {
		for _, value := range values {
			if key == "Authorization" {
				tokenString = strings.Replace(value, "Bearer ", "", -1)

			}
		}
	}

	// decode token with secret code bermaslaah
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("bermaslaah"), nil
	})

	if !token.Valid {
		fmt.Println("invalid jwt token!!!")
	}

	dataUpdatePerformanceID := bson.M{"_id": req.ID}

	var reqPerformance model.DataPerformance
	filter := bson.M{"_id": req.ID}
	errUser := cr.mongoDB.Collection("performances").FindOne(ctx, filter).Decode(&reqPerformance)

	if errUser != nil {
		log.Println("error")
	}

	RevisionCount := reqPerformance.RevisionCount
	if req.Status == "review" {
		RevisionCount += 1
	}

	dataObjectID := bson.M{"$set": bson.M{
		"status":         req.Status,
		"update_at":      time.Now(),
		"update_by":      claims["id"].(string),
		"revision_count": RevisionCount,
	}}
	_, err = cr.mongoDB.Collection("performances").UpdateOne(ctx, dataUpdatePerformanceID, dataObjectID)
	if err != nil {
		log.Println("error")
	}

	return err
}

func (cr PerformanceRepository) DeleteDataPerformance(c echo.Context, ctx context.Context, req model.DataPerformanceRequest) (err error) {
	tokenString := ""

	// get jwt bearer token from context
	for key, values := range c.Request().Header {
		for _, value := range values {
			if key == "Authorization" {
				tokenString = strings.Replace(value, "Bearer ", "", -1)

			}
		}
	}

	// decode token with secret code bermaslaah
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("bermaslaah"), nil
	})

	if !token.Valid {
		fmt.Println("invalid jwt token!!!")
	}

	dataUpdatePerformanceID := bson.M{"_id": req.ID}

	dataObjectID := bson.M{"$set": bson.M{
		"deleted_at": time.Now(),
		"deleted_by": claims["id"].(string),
		"deleted":    true,
	}}

	_, err = cr.mongoDB.Collection("performances").UpdateOne(ctx, dataUpdatePerformanceID, dataObjectID)
	if err != nil {
		log.Println("error")
	}

	return err
}

func sendMail(to []string, cc []string, subject, message string) error {
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Info KPI Indochat <infokpiindochat@gmail.com>"
const CONFIG_AUTH_EMAIL = "infokpiindochat@gmail.com"
const CONFIG_AUTH_PASSWORD = "@Password1"
