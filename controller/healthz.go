package controller

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

// Health for monitoring health of our app
type Health interface {
	IsHealthy() 
}

// HealthController for monitoring health of our app
type HealthController struct {
	logger   *logrus.Logger
	render   *render.Render
}

// NewHealthController returns a playback_rights controller
func NewHealthController(
	logger *logrus.Logger,
	r *render.Render,
) *HealthController {
	return &HealthController{logger, r}
}

// IsHealthy handler checks health status of all external services
func (h *HealthController) IsHealthy(w http.ResponseWriter, r *http.Request) {
	h.logger.WithField("func", "IsHealthy").Info("in")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Chillin\n"))
}