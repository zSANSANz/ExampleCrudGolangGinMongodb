module chatnews-api

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/gobeam/mongo-go-pagination v0.0.8
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/labstack/echo/v4 v4.6.3
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.10.1
	github.com/swaggo/echo-swagger v1.1.4
	go.mongodb.org/mongo-driver v1.8.2
	go.uber.org/zap v1.20.0
	golang.org/x/crypto v0.0.0-20211108221036-ceb1ce70b4fa
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)
docker build -t kpi-backend-img .
  207  sudo docker build -t kpi-backend-img .
  216  sudo docker run --name="kpi-backend-container" -d -p 9595:9595 --restart always kpi-backend-img
