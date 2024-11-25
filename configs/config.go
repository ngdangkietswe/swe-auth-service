package configs

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

var (
	GlobalConfig = &Configuration{}
)

func init() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Can't read config file: %v", err)
		return
	}

	config := &Configuration{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("Can't unmarshal config: %v", err)
		return
	}

	GlobalConfig = config

	return
}

type Configuration struct {
	AppName      string        `mapstructure:"APP_NAME"`
	GrpcPort     int           `mapstructure:"GRPC_PORT"`
	DbHost       string        `mapstructure:"DB_HOST"`
	DbPort       int           `mapstructure:"DB_PORT"`
	DbUser       string        `mapstructure:"DB_USER"`
	DbPassword   string        `mapstructure:"DB_PASSWORD"`
	DbName       string        `mapstructure:"DB_NAME"`
	JwtSecret    string        `mapstructure:"JWT_SECRET"`
	JwtIssuer    string        `mapstructure:"JWT_ISSUER"`
	JwtExp       time.Duration `mapstructure:"JWT_EXPIRATION"`
	KafkaBrokers string        `mapstructure:"KAFKA_BROKER"`
}
