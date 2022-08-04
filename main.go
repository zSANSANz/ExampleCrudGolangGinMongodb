package main

import (
	"fmt"
	"log"

	httpDelivery "chatnews-api/api/delivery/http"
	"chatnews-api/api/model"
	repository "chatnews-api/api/repository/impl"
	usecase "chatnews-api/api/usecase/impl"
	"chatnews-api/lib/config"
	"chatnews-api/lib/db"
	"chatnews-api/lib/logging"
	"chatnews-api/middleware/controller"
	"chatnews-api/middleware/routes"
	"chatnews-api/middleware/security"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"

	JwtConfig "chatnews-api/middleware/config"
	JwtRepository "chatnews-api/middleware/repository"
)

var userControllerAuth *controller.UserController
var performanceControllerAuth *controller.PerformanceController

//to initialize viper config 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
func init() {
	config.SetConfigFile("config", "lib/config", "json")
}

func main() {
	envConfig := getConfig()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowHeaders:     []string{"*"},
		// AllowCredentials: true,
		// AllowOrigins:     []string{"*"},
		// AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.Use(logging.MiddlewareLogging)

	// Mongo
	mongo, err := db.Connect(envConfig.Mongo)
	if err != nil {
		log.Println(err)
		return
	}

	// category
	categoryRepo := repository.NewCategoryRepository(mongo)
	categoryUseCase := usecase.NewCategoryUseCase(&envConfig, categoryRepo)

	httpDelivery.CategoryRouter(e, categoryUseCase)

	// employee
	employeeRepo := repository.NewEmployeeRepository(mongo)
	employeeUseCase := usecase.NewEmployeeUseCase(&envConfig, employeeRepo)

	httpDelivery.EmployeeRouter(e, employeeUseCase)

	// head devision
	headDevisionRepo := repository.NewHeadDevisionRepository(mongo)
	headDevisionUseCase := usecase.NewHeadDevisionUseCase(&envConfig, headDevisionRepo)

	httpDelivery.HeadDevisionRouter(e, headDevisionUseCase)

	// performance
	performanceRepo := repository.NewPerformanceRepository(mongo)
	performanceUseCase := usecase.NewPerformanceUseCase(&envConfig, performanceRepo)

	httpDelivery.PerformanceRouter(e, performanceUseCase)

	// jwt auth
	mongoConnection, errorMongoConn := JwtConfig.MongoConnection()

	if errorMongoConn != nil {
		log.Println("Error when connect mongo : ", errorMongoConn.Error())
	}

	userRepositoryAuth := JwtRepository.NewUserRepository(mongoConnection)
	authValidator := security.NewAuthValidator(userRepositoryAuth)
	userControllerAuth = controller.NewUserController(userRepositoryAuth, authValidator)
	routes.GetUserApiRoutes(e, userControllerAuth)

	performanceRepositoryAuth := JwtRepository.NewPerformanceRepository(mongoConnection)
	performanceControllerAuth = controller.NewPerformanceController(performanceRepositoryAuth, authValidator)
	routes.GetPerformanceApiRoutes(e, performanceControllerAuth)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s%s%v", envConfig.Host, ":", envConfig.Port)))
}

func getConfig() model.EnvConfig {
	return model.EnvConfig{
		Host: config.GetString("host.address"),
		Port: config.GetInt("host.port"),
		Mongo: db.MongoConfig{
			Timeout:  config.GetInt("database_" + config.GetString("env") + ".mongodb.timeout"),
			DBname:   config.GetString("database_" + config.GetString("env") + ".mongodb.dbname"),
			Username: config.GetString("database_" + config.GetString("env") + ".mongodb.user"),
			Password: config.GetString("database_" + config.GetString("env") + ".mongodb.password"),
			Host:     config.GetString("database_" + config.GetString("env") + ".mongodb.host"),
			Port:     config.GetString("database_" + config.GetString("env") + ".mongodb.port"),
		},
	}
}
