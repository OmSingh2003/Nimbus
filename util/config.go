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
	viper.SetDefault("ACCESS_TOKEN_DURATION", "15m")
	viper.SetDefault("REFRESH_TOKEN_DURATION", "24h")
	viper.SetDefault("EMAIL_SENDER_NAME", "Nimbus")
	// Note: DB_SOURCE, REDIS_ADDRESS, TOKEN_SYMMETRIC_KEY, EMAIL credentials should come from env vars


	err = viper.Unmarshal(&config)
	return
}
