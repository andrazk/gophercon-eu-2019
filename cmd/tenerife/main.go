package main

import (
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"tenerife/internal/handler"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	r := mux.NewRouter()
	r.HandleFunc("/hola", http.HandlerFunc(handler.HolaHandler))

	server := http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: r,
	}

	server.ListenAndServe()
}
