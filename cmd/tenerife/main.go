package main

import (
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"tenerife/internal/handler"
)

func main() {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{})

	logger.Info("Starting Tenerife app")

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if port == "" {
		logger.Fatal("Port is not defined")
	}

	r := mux.NewRouter()
	r.HandleFunc("/hola", http.HandlerFunc(handler.HolaHandler(logger)))

	server := http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: r,
	}

	server.ListenAndServe()
}
