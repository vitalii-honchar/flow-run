package flowrun

import (
	"context"
	"flow-run/internal/flowrun/config"
	"flow-run/internal/flowrun/infra/api"
	"flow-run/internal/flowrun/infra/api/handler/health"
	"flow-run/internal/flowrun/infra/api/middleware"
	"flow-run/internal/flowrun/infra/database"
	"flow-run/internal/lib/logger"
)

type FlowRun struct {
	Config *config.Config
	DB     *database.Database

	components []component
}

type component struct {
	name  string
	start func(ctx context.Context) error
	stop  func(ctx context.Context) error
}

func NewFlowRun() (*FlowRun, error) {
	cfg, err := config.FromEnv()
	if err != nil {
		return nil, err
	}

	db, err := database.NewDatabase(cfg.DatabaseConfig)
	if err != nil {
		return nil, err
	}

	server := api.NewServer(
		[]api.Middleware{
			middleware.NewLoggingMiddleware(),
		},
		[]api.Handler{
			health.NewHealthHandler(db),
		},
		cfg,
	)

	return &FlowRun{
		Config: cfg,
		DB:     db,
		components: []component{
			{name: "database", stop: db.Stop},
			{name: "server", start: server.Start, stop: server.Stop},
		},
	}, nil
}

func (fr *FlowRun) Start(ctx context.Context) error {
	logger.Log.Info("Starting FlowRun server")

	for _, component := range fr.components {
		if component.start != nil {
			if err := component.start(ctx); err != nil {
				logger.Log.WithError(err).Warnf("Failed to start component %s", component.name)
				return err
			}
		}
	}

	return nil
}

func (fr *FlowRun) Stop(ctx context.Context) error {
	logger.Log.Info("Stopping FlowRun server")

	for _, component := range fr.components {
		if component.stop != nil {
			if err := component.stop(ctx); err != nil {
				logger.Log.WithError(err).Warnf("Failed to stop component %s", component.name)
			}
		}
	}
	logger.Log.Info("Stopped FlowRun server")

	return nil
}
