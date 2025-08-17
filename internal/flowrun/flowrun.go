package flowrun

import (
	"context"
	"flow-run/internal/flowrun/config"
	"flow-run/internal/lib/logger"
	"fmt"
	"net/http"
	"time"
)

type FlowRun struct {
	Config *config.Config
	server *http.Server
}

func NewFlowRun() (*FlowRun, error) {
	logger.Log.Info("Starting FlowRun")

	cfg, err := config.FromEnv()
	if err != nil {
		return nil, err
	}

	logger.Log.Info("FlowRun configuration loaded")

	return &FlowRun{
		Config: cfg,
	}, nil
}

func (fr *FlowRun) Start() error {
	logger.Log.Info("Starting FlowRun server")

	// Create HTTP server
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	fr.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", fr.Config.ServerHost, fr.Config.ServerPort),
		Handler: mux,
	}

	logger.Log.WithField("address", fr.server.Addr).Info("FlowRun server started")

	// Start server in a goroutine so it doesn't block
	go func() {
		if err := fr.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.WithError(err).Error("Failed to start server")
		}
	}()

	return nil
}

func (fr *FlowRun) Stop() error {
	logger.Log.Info("Stopping FlowRun server")

	if fr.server == nil {
		return nil
	}

	// Create a timeout context for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Gracefully shutdown the server
	if err := fr.server.Shutdown(ctx); err != nil {
		logger.Log.WithError(err).Error("Failed to gracefully shutdown server")
		return err
	}

	logger.Log.Info("FlowRun server stopped")
	return nil
}
