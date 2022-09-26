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

var ConfigGlobal Config

func InitConfig() Config {
	// ConfigGlobal = Config{
	// 	SERVER_ADDRESS: GetOrDefault("SERVER_ADDRESS", "0.0.0.0:8888"),
	// 	DB_USERNAME:    GetOrDefault("DB_USERNAME", "root"),
	// 	DB_PASSWORD:    GetOrDefault("DB_PASSWORD", ""),
	// 	DB_NAME:        GetOrDefault("DB_NAME", "generate_invoices"),
	// 	DB_PORT:        GetOrDefault("DB_PORT", "3306"),
	// 	DB_HOST:        GetOrDefault("DB_HOST", "127.0.0.1"),
	// 	JWT_KEY:        GetOrDefault("JWT_KEY", "AlphaWolf"),
	// }
	ConfigGlobal = Config{
		SERVER_ADDRESS: GetOrDefault("SERVER_ADDRESS", "0.0.0.0:8888"),
		DB_USERNAME:    GetOrDefault("DB_USERNAME", "invoice-backend"),
		DB_PASSWORD:    GetOrDefault("DB_PASSWORD", "35GenerateInvoices"),
		DB_NAME:        GetOrDefault("DB_NAME", "generate_invoices"),
		DB_PORT:        GetOrDefault("DB_PORT", "3306"),
		DB_HOST:        GetOrDefault("DB_HOST", "localhost"),
		JWT_KEY:        GetOrDefault("JWT_KEY", "AlphaWolf"),
	}
	return ConfigGlobal
}

func GetOrDefault(envName string, defaultValue string) string {
	if value, ok := os.LookupEnv(envName); ok {
		return value
	}

	return defaultValue
}
