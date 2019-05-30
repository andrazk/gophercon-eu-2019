package main

import (
	"context"
	"flag"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/peterbourgon/ff"
	"github.com/sirupsen/logrus"

	"tenerife/internal/application"
	"tenerife/internal/diagnostics"
)

func main() {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)

	logger.Info("Starting Tenerife app")
	logger.Infof("Version %s", diagnostics.Version())

	fs := flag.NewFlagSet("tenerife", flag.ExitOnError)
	var (
		host            = fs.String("host", "localhost", "server host")
		port            = fs.String("port", "8080", "server port")
		portDiagnostics = fs.String("diagnostics-port", "9090", "diagnostics port")
	)

	ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarPrefix("TENERIFE"),
	)

	// Config validation
	if *port == "" {
		logger.Fatal("Port is not defined")
	}

	r := mux.NewRouter()
	r.HandleFunc("/hola", http.HandlerFunc(application.HolaHandler(logger)))

	// We want to consume all signal in cached channel
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	shutdown := make(chan error, 1)

	var server http.Server
	{
		addr := net.JoinHostPort(*host, *port)
		server = http.Server{
			Addr:    addr,
			Handler: r,
		}
	}

	var serverDiagnostics http.Server
	{
		addr := net.JoinHostPort(*host, *portDiagnostics)
		serverDiagnostics = http.Server{
			Addr:    addr,
			Handler: diagnosticsMux(logger),
		}
	}

	go func() {
		logger.Infof("Starting Application server. Listening on %s", server.Addr)
		err := server.ListenAndServe()
		shutdown <- err
	}()

	go func() {
		logger.Infof("Starting Diagnostics server. Listening on %s", serverDiagnostics.Addr)
		err := serverDiagnostics.ListenAndServe()
		shutdown <- err
	}()

	select {
	case killSignal := <-interrupt:
		switch killSignal {
		case os.Interrupt:
			logger.Print("Got SIGINT...")
		case syscall.SIGTERM:
			logger.Print("Got SIGTERM...")
		}
	case err := <-shutdown:
		logger.Errorf("Got an error %v", err)
	}

	logger.Info("Application server is stopping")
	err := server.Shutdown(context.Background())
	if err != nil {
		logger.Fatalf("Application shutdown failed with error %v", err)
	}

	logger.Info("Diagnostics server is stopping")
	err = serverDiagnostics.Shutdown(context.Background())
	if err != nil {
		logger.Fatalf("Diagnostics shutdown failed with error %v", err)
	}
}

func diagnosticsMux(logger *logrus.Logger) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/healthz", http.HandlerFunc(diagnostics.LivelinessHandler(logger)))
	r.HandleFunc("/readyz", http.HandlerFunc(diagnostics.ReadinessHandler(logger)))

	return r
}

func server(host, port string, handler http.Handler) http.Server {
	return http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: handler,
	}
}
