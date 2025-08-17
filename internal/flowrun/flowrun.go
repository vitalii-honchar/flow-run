package flowrun

import (
	"context"
	"flow-run/internal/flowrun/config"
	"flow-run/internal/flowrun/infra/database"
	"flow-run/internal/lib/logger"
)

type FlowRun struct {
	Config *config.Config
	DB     *database.Database

	components []component
}

type component struct {
	name string
	stop func(ctx context.Context) error
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

	return &FlowRun{
		Config: cfg,
		DB:     db,
		components: []component{
			{name: "database", stop: db.Stop},
		},
	}, nil
}

func (fr *FlowRun) Start() error {
	logger.Log.Info("Starting FlowRun server")

	return nil
}

func (fr *FlowRun) Stop(ctx context.Context) error {
	logger.Log.Info("Stopping FlowRun server")

	for _, component := range fr.components {
		if err := component.stop(ctx); err != nil {
			logger.Log.WithError(err).Warnf("Failed to stop component %s", component.name)
		}
	}
	logger.Log.Info("Stopped FlowRun server")

	return nil
}
