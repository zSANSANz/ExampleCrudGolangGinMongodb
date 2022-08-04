package util

import (
	"chatnews-api/middleware/config"
	"chatnews-api/middleware/model"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GenerateJwtToken(user *model.User) (string, error) {
	expTimeMs, _ := strconv.Atoi(config.JWTExpirationMs)
	exp := time.Now().Add(time.Millisecond * time.Duration(expTimeMs)).Unix()
	name := fmt.Sprintf("%s %s", user.FirstName, user.LastName)

	// Set custom claims
	claims := &model.JwtCustomClaims{
		user.ID.Hex(),
		name,
		jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response
	jwt, err := token.SignedString([]byte(config.JWTSecret))
	return jwt, err

}

func GetUserIdFromToken(c echo.Context) string {
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

	return claims["id"].(string)

}
