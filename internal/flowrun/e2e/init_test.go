package e2e

import (
	"context"
	"flow-run/internal/flowrun"
	"flow-run/internal/lib/logger"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

var (
	testFlowRun *flowrun.FlowRun
)

const (
	startTimeout = 60 * time.Second
	stopTimeout  = 60 * time.Second
)

func TestMain(m *testing.M) {
	envPath := filepath.Join(".", ".env")
	if err := godotenv.Load(envPath); err != nil {
		envPath = filepath.Join("..", "..", "..", ".env")
		_ = godotenv.Load(envPath)
	}

	var err error
	testFlowRun, err = flowrun.NewFlowRun()
	if err != nil {
		logger.WithError(err).Fatal("Failed to create FlowRun instance")
	}

	ctx, cancel := context.WithTimeout(context.Background(), startTimeout)
	defer cancel()

	if err := testFlowRun.Start(ctx); err != nil {
		logger.WithError(err).Fatal("Failed to start FlowRun instance")
	}
	defer stopFlowRun()

	code := m.Run()

	os.Exit(code)
}

func stopFlowRun() {
	if testFlowRun != nil {
		ctx, cancel := context.WithTimeout(context.Background(), stopTimeout)
		defer cancel()

		err := testFlowRun.Stop(ctx)
		if err != nil {
			logger.WithError(err).Error("Failed to stop FlowRun instance")
		}
	}
}
