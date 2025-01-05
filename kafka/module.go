package kafka

import (
	"github.com/ngdangkietswe/swe-auth-service/kafka/producer"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	producer.NewKProducer,
)
