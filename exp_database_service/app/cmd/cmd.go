package cmd

import (
	"app/internal/logger"
	"app/internal/service"

	"context"
	"os"

	"go.uber.org/zap"
)

func RunService() {
	var log *zap.SugaredLogger = logger.Logger(os.Getenv("LOG_LEVEL"))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	svc := service.New(log, ctx)
	svc.Run()
}
