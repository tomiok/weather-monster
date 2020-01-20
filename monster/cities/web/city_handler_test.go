package web

import (
	"github.com/go-chi/chi"
	"github.com/tomiok/weather-monster/internal/logs"
	"net/http"
	"net/http/httptest"
	"testing"
)

func router() *chi.Mux {
	return chi.NewRouter()
}

func Test_CityHandler_DeleteCity(t *testing.T) {
	logs.InitDefault("test")
	req, err := http.NewRequest(http.MethodDelete, "/cities/BAD-ID", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	var h CityHandler
	r := router()
	r.Delete("/cities/BAD-ID", h.DeleteCityHandler)
	r.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func Test_CityHandler_UpdateCity(t *testing.T) {
	logs.InitDefault("test")
	req, err := http.NewRequest(http.MethodPatch, "/cities/BAD-ID", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	var h CityHandler
	r := router()
	r.Patch("/cities/BAD-ID", h.UpdateCityHandler)
	r.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
