package logger

import "go.uber.org/fx"

var Module = fx.Provide(NewZapLogger)
