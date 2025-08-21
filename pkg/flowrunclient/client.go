package flowrunclient

import (
	"context"
	"encoding/json"
	"flow-run/pkg/flowrunclient/model"
	"fmt"
	"io"
	"net/http"
)

type FlowRunClient interface {
	GetHealth(ctx context.Context) (*model.HealthResponse, error)
}

type flowRunClient struct {
	baseURL string
}

func NewFlowRunClient(baseURL string) FlowRunClient {
	return &flowRunClient{
		baseURL: baseURL,
	}
}

func (c *flowRunClient) GetHealth(ctx context.Context) (*model.HealthResponse, error) {
	return get[model.HealthResponse](c.baseURL, "/v1/health")
}

func get[T any](baseURL string, endpoint string) (*T, error) {
	resp, err := http.Get(baseURL + endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result T
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
