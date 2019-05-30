package application

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// HolaHandler writes Hola in response
func HolaHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		logger.Info("Say hello in spanish")
		w.Write([]byte("Hola!"))
	}
}
