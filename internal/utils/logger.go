package utils

import (
	"context"

	"github.com/rs/zerolog"
)

func GetLoggerFromCtx(ctx context.Context, fallback *zerolog.Logger) *zerolog.Logger {
	logger, ok := ctx.Value("logger").(*zerolog.Logger)

	if ok {
		return logger
	}

	fallback.Error().Msg("fail to get logger from context")
	return fallback
}
