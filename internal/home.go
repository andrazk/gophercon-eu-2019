package internal

import "net/http"

// HolaHandler writes Hola in response
func HolaHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hola!"))
}
