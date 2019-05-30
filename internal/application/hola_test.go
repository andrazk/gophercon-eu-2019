package application_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"tenerife/internal/application"

	"github.com/sirupsen/logrus"
)

func TestHolaHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hola", nil)
	rr := httptest.NewRecorder()

	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)

	application.HolaHandler(logger)(rr, req)
	if got, want := rr.Result().StatusCode, http.StatusOK; want != got {
		t.Errorf("Expected %d, got %d", want, got)
	}

}
