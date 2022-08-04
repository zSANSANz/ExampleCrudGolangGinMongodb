package config

import "os"

var (
	ServerPort      = GetEnv("SERVER_PORT", "9595")
	MongoUrl        = GetEnv("MONGODB_URL", "mongodb://mongoAdmin:blabla@172.0.0.1:27017")
	MongoDatabase   = GetEnv("MONGODB_DATABASE", "kpi_production")
	JWTSecret       = GetEnv("JWT_SECRET", "bermaslaah")
	JWTExpirationMs = GetEnv("JWT_EXPIRATION_MS", "86400000")
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
