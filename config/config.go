package config

import "os"

type Config struct {
	SERVER_ADDRESS string
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_PORT        string
	DB_HOST        string
	DB_NAME        string
	JWT_KEY        string
}

func InitConfig() Config {
	return Config{
		SERVER_ADDRESS: GetOrDefault("SERVER_ADDRESS", "0.0.0.0:8888"),
		DB_USERNAME:    GetOrDefault("DB_USERNAME", "admin"),
		DB_PASSWORD:    GetOrDefault("DB_PASSWORD", "admin12345"),
		DB_NAME:        GetOrDefault("DB_NAME", "generate_invoices"),
		DB_PORT:        GetOrDefault("DB_PORT", "3306"),
		DB_HOST:        GetOrDefault("DB_HOST", "database-generate-invoices.cv8cnjvlnjwz.us-west-1.rds.amazonaws.com"),
		JWT_KEY:        GetOrDefault("JWT_KEY", "AlphaWolf"),
	}
}

func GetOrDefault(envName string, defaultValue string) string {
	if value, ok := os.LookupEnv(envName); ok {
		return value
	}

	return defaultValue
}
