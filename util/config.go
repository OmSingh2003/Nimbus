package util

import (
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
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// Set defaults first
	viper.SetDefault("DB_DRIVER", "postgres")
	viper.SetDefault("HTTP_SERVER_ADDRESS", "0.0.0.0:8080")
	viper.SetDefault("GRPC_SERVER_ADDRESS", "0.0.0.0:9090")
	viper.SetDefault("REDIS_ADDRESS", "redis://default:QyjsX69AHhN7RXTcdAMXv2G2Ow4CqOFp@redis-12951.c44.us-east-1-2.ec2.redns.redis-cloud.com:12951")
	viper.SetDefault("TOKEN_SYMMETRIC_KEY", "12345678901234567890123456789012")
	viper.SetDefault("ACCESS_TOKEN_DURATION", "15m")
	viper.SetDefault("REFRESH_TOKEN_DURATION", "24h")
	viper.SetDefault("EMAIL_SENDER_NAME", "Simple bank")

	// Try to read config file, but don't fail if it doesn't exist (for production)
	err = viper.ReadInConfig()
	if err != nil {
		// If config file doesn't exist, we'll use environment variables only
		// This is expected in production environments like Render
		// Clear the error since we can work without the config file
		err = nil
	}

	err = viper.Unmarshal(&config)
	return
}
