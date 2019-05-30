package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"tenerife/internal/handler"
)

func TestHolaHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hola", nil)
	rr := httptest.NewRecorder()

	handler.HolaHandler(rr, req)
	if got, want := rr.Result().StatusCode, http.StatusOK; want != got {
		t.Errorf("Expected %d, got %d", want, got)
	}
}
