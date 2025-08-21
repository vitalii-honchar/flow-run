package main

import (
	"context"
	"flow-run/internal/flowrun"
	"flow-run/internal/lib/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	startTimeout = 60 * time.Second
	stopTimeout  = 60 * time.Second
)

func main() {
	fr, err := flowrun.NewFlowRun()
	if err != nil {
		logger.Log.WithError(err).Fatal("Failed to create FlowRun instance")
	}

	ctx, cancel := context.WithTimeout(context.Background(), startTimeout)
	defer cancel()

	if err := fr.Start(ctx); err != nil {
		logger.Log.WithError(err).Fatal("Failed to start FlowRun server")
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	logger.Log.WithField("signal", sig.String()).Info("Received shutdown signal")

	ctx, cancel = context.WithTimeout(context.Background(), stopTimeout)
	defer cancel()

	if err := fr.Stop(ctx); err != nil {
		logger.Log.WithError(err).Error("Failed to stop FlowRun server gracefully")
		os.Exit(1)
	}

	logger.Log.Info("FlowRun shutdown complete")
}
