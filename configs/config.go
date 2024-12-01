package configs

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
	"time"
)

var (
	GlobalConfig = &Configuration{}
)

func init() {
	env := os.Getenv("K8S_ENV")
	log.Printf("K8S_ENV is set to %s", env)
	if strings.ToLower(env) == "prod" {
		log.Println("Using production config")
		viper.AutomaticEnv()
	} else {
		log.Println("Using local config")
		viper.AddConfigPath("./configs")
		viper.SetConfigName("config")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Can't read config file: %v", err)
			return
		}
	}

	err := viper.Unmarshal(&GlobalConfig)
	if err != nil {
		log.Fatalf("Can't unmarshal config: %v", err)
	}
}

type Configuration struct {
	AppName         string        `mapstructure:"APP_NAME"`
	GrpcPort        int           `mapstructure:"GRPC_PORT"`
	DbHost          string        `mapstructure:"DB_HOST"`
	DbPort          int           `mapstructure:"DB_PORT"`
	DbUser          string        `mapstructure:"DB_USER"`
	DbPassword      string        `mapstructure:"DB_PASSWORD"`
	DbName          string        `mapstructure:"DB_NAME"`
	JwtSecret       string        `mapstructure:"JWT_SECRET"`
	JwtIssuer       string        `mapstructure:"JWT_ISSUER"`
	JwtExp          time.Duration `mapstructure:"JWT_EXPIRATION"`
	RefreshTokenExp time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRATION"`
	KafkaBrokers    string        `mapstructure:"KAFKA_BROKER"`
	RedisHost       string        `mapstructure:"REDIS_HOST"`
	RedisPort       int           `mapstructure:"REDIS_PORT"`
}
