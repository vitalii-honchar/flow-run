package health

import (
	"context"
	"flow-run/internal/lib/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

const healthGroup = "v1/health"

type (
	HealthHandler struct {
		dbPinger dbPinger
	}

	dbPinger interface {
		Ping(ctx context.Context) error
	}
)

func NewHealthHandler(dbPinger dbPinger) *HealthHandler {
	return &HealthHandler{
		dbPinger: dbPinger,
	}
}

func (h *HealthHandler) Group() string {
	return healthGroup
}

func (h *HealthHandler) Method() string {
	return http.MethodGet
}

func (h *HealthHandler) Path() string {
	return "/"
}

func (h *HealthHandler) Handle(c *gin.Context) {

	if err := h.dbPinger.Ping(c.Request.Context()); err != nil {
		logger.WithError(err).Error("Database ping failed")
		c.JSON(http.StatusInternalServerError, gin.H{"status": "unhealthy"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
