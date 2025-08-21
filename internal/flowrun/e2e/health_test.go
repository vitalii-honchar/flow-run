package e2e

import (
	"context"
	"flow-run/pkg/flowrunclient/model"
	"testing"
	"time"
)

func TestServiceStarted(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := testClient.GetHealth(ctx)
	if err != nil {
		t.Fatalf("Failed to get health: %v", err)
	}

	if response == nil {
		t.Fatal("Health response is nil")
	}

	if response.Status != model.HealthStatusUp {
		t.Errorf("Expected health status to be %s, got %s", model.HealthStatusUp, response.Status)
	}
}
