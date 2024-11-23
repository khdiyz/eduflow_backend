package config

import (
	"eduflow/pkg/logger"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	HTTPHost string
	HTTPPort int

	Environment string
	Debug       bool

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

	RedisHost     string
	RedisPort     int
	RedisPassword string

	JWTSecret                string
	JWTAccessExpirationHours int
	JWTRefreshExpirationDays int

	HashKey string

	MinioEndpoint    string
	MinioAccessKeyId string
	MinioSecretKey   string
	MinioBucketName  string
	MinioUseSSL      bool

	SuperAdminUsername string
	SuperAdminPassword string
}

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{
			HTTPHost:    cast.ToString(getOrReturnDefault("HOST", "localhost")),
			HTTPPort:    cast.ToInt(getOrReturnDefault("PORT", 4040)),
			Environment: cast.ToString(getOrReturnDefault("ENVIRONMENT", EnvironmentDevelopment)),
			Debug:       cast.ToBool(getOrReturnDefault("DEBUG", true)),

			PostgresHost:     cast.ToString(getOrReturnDefault("POSTGRES_HOST", "142.93.102.185")),
			PostgresPort:     cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432)),
			PostgresDatabase: cast.ToString(getOrReturnDefault("POSTGRES_DB", "eduflow_test_db")),
			PostgresUser:     cast.ToString(getOrReturnDefault("POSTGRES_USER", "khdiyz")),
			PostgresPassword: cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "password")),

			RedisHost:     cast.ToString(getOrReturnDefault("REDIS_HOST", "142.93.102.185")),
			RedisPort:     cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379)),
			RedisPassword: cast.ToString(getOrReturnDefault("REDIS_PASSWORD", "")),

			JWTSecret:                cast.ToString(getOrReturnDefault("JWT_SECRET", "eduflow-forever")),
			JWTAccessExpirationHours: cast.ToInt(getOrReturnDefault("JWT_ACCESS_EXPIRATION_HOURS", "12")),
			JWTRefreshExpirationDays: cast.ToInt(getOrReturnDefault("JWT_REFRESH_EXPIRATION_DAYS", "3")),

			HashKey: cast.ToString(getOrReturnDefault("HASH_KEY", "skd32r8$wdahHSdqw")),

			MinioEndpoint:    cast.ToString(getOrReturnDefault("MINIO_ENDPOINT", "142.93.102.185:9000")),
			MinioAccessKeyId: cast.ToString(getOrReturnDefault("MINIO_ACCESS_KEY_ID", "b5qxOurcZQuzJqcztqTR")),
			MinioSecretKey:   cast.ToString(getOrReturnDefault("MINIO_SECRET_KEY", "ylGnuSIiervvaUN9MVKRgDj2aEC3Tru7WSEdeSOx")),
			MinioBucketName:  cast.ToString(getOrReturnDefault("MINIO_BUCKET_NAME", "eduflow")),
			MinioUseSSL:      cast.ToBool(getOrReturnDefault("MINIO_USE_SLL", false)),

			SuperAdminUsername: cast.ToString(getOrReturnDefault("SUPER_ADMIN_USERNAME", "")),
			SuperAdminPassword: cast.ToString(getOrReturnDefault("SUPER_ADMIN_PASSWORD", "")),
		}
	})

	return instance
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	err := godotenv.Load(".env")
	if err != nil {
		logger.GetLogger().Error(err)
	}
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
