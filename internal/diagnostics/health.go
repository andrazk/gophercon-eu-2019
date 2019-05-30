package diagnostics

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// LivelinessHandler for k8s health probe
func LivelinessHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		logger.Info("Liveliness probe")
		w.WriteHeader(http.StatusOK)
	}
}
