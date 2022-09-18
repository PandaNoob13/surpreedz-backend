package config

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
)

type ApiConfig struct {
	Url string
}

type DbConfig struct {
	DataSourceName string
}

// type GrpcConfig struct {
// 	UrlGrpc string
// }

type AzureConfig struct {
	AccountName string
	AccountKey  string
	ServiceUrl  string
}

type TokenConfig struct {
	ApplicationName     string
	JwtSingingMethod    *jwt.SigningMethodHMAC
	JwtSignatureKey     string
	AccessTokenLifeTIme time.Duration
	Client              *redis.Client
}

type Config struct {
	ApiConfig
	DbConfig
	TokenConfig
	AzureConfig
}

func (c *Config) readConfig() {
	// api := os.Getenv("API_URL")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")

	api := "localhost:8080"
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := ""
	dbName := ""

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	c.ApiConfig = ApiConfig{Url: api}
	c.DbConfig = DbConfig{DataSourceName: dsn}
	// c.GrpcConfig = GrpcConfig{UrlGrpc: grpcUrl}
	c.TokenConfig = TokenConfig{
		ApplicationName:     "SURPREEDZ",
		JwtSingingMethod:    jwt.SigningMethodHS256,
		JwtSignatureKey:     "5URPR33DZ",
		AccessTokenLifeTIme: 6000 * time.Second,
		Client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
	c.AzureConfig = AzureConfig{
		AccountName: "surpreedzstorage",
		AccountKey:  "mVRAySXmfDDTqMoVi12+6bNie/4xoxrkouZYsEGkadlzzYnGRVjLu+lzbvx6x/krZdgvFVB69+9u+AStBDi0Ow==",
		ServiceUrl:  "https://surpreedzstorage.blob.core.windows.net/",
	}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
