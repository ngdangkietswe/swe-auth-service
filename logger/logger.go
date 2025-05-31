package logger

import (
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
)

func NewZapLogger() (*logger.Logger, error) {
	instance, err := logger.NewLogger(
		config.GetString("APP_NAME", "swe-auth"),
		config.GetString("APP_ENV", "dev"),
		"debug",
		"",
	)

	if err != nil {
		return nil, err
	}

	return instance, nil
}
