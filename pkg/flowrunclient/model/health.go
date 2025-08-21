package model

type HealthStatus string

const (
	HealthStatusUp   HealthStatus = "up"
	HealthStatusDown HealthStatus = "down"
)

type HealthResponse struct {
	Status HealthStatus `json:"status"`
}

func NewHealthResponse(status HealthStatus) *HealthResponse {
	return &HealthResponse{
		Status: status,
	}
}
