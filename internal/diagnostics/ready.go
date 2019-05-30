package diagnostics

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// ReadinessHandler for k8s health probe
func ReadinessHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		logger.Info("Readiness probe")
		w.WriteHeader(http.StatusOK)
	}
}
