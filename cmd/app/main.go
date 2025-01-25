package main

import (
	"context"

	"go.uber.org/zap"

	a "github.com/blckvia/go-db/internal/app"
)

func main() {
	ctx := context.Background()

	logger := zap.Must(zap.NewProduction())
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Error("failed to sync logger", zap.Error(err))
		}
	}(logger)

	_ = a.NewApp(ctx, logger)
}
