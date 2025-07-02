package util

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

// Config stores all the configuration of the application
// The values are read by viper from a config file or environment variables
type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	RedisAddress         string        `mapstructure:"REDIS_ADDRESS"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
	EmailVerificationURL string        `mapstructure:"EMAIL_VERIFICATION_URL"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	// Read environment variables directly using os.Getenv for critical values
	config.DBDriver = getEnvOrDefault("DB_DRIVER", "postgres")
	config.DBSource = getEnvOrDefault("DB_SOURCE", "")
	config.RedisAddress = getEnvOrDefault("REDIS_ADDRESS", "")
	config.HTTPServerAddress = getEnvOrDefault("HTTP_SERVER_ADDRESS", "0.0.0.0:8080")
	config.GRPCServerAddress = getEnvOrDefault("GRPC_SERVER_ADDRESS", "0.0.0.0:9090")
	config.TokenSymmetricKey = getEnvOrDefault("TOKEN_SYMMETRIC_KEY", "")
	config.EmailSenderName = getEnvOrDefault("EMAIL_SENDER_NAME", "Nimbus")
	config.EmailSenderAddress = getEnvOrDefault("EMAIL_SENDER_ADDRESS", "")
	config.EmailSenderPassword = getEnvOrDefault("EMAIL_SENDER_PASSWORD", "")
	config.EmailVerificationURL = getEnvOrDefault("EMAIL_VERIFICATION_URL", "")

	// Parse duration values
	accessTokenDuration := getEnvOrDefault("ACCESS_TOKEN_DURATION", "15m")
	config.AccessTokenDuration, err = time.ParseDuration(accessTokenDuration)
	if err != nil {
		return config, err
	}

	refreshTokenDuration := getEnvOrDefault("REFRESH_TOKEN_DURATION", "24h")
	config.RefreshTokenDuration, err = time.ParseDuration(refreshTokenDuration)
	if err != nil {
		return config, err
	}

	return config, nil
}

// getEnvOrDefault gets an environment variable or returns a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
