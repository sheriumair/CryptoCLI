package e2e

import (
	"backend/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test the status code 200 of About endpoint
func TestHandleAboutCommand(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handlers.HandleAboutCommand))
	defer ts.Close()
	client := &http.Client{}
	resp, err := client.Get(ts.URL + "/about")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

// Test the status code 200 of Help endpoint
func TestHandleHelpCommand(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(handlers.HandleHelpCommand))
	defer srv.Close()
	client := srv.Client()
	resp, err := client.Get(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}
}

// Test the status code 200 of FetchPrice endpoint as currency entered is valid
func TestHandleFetchPriceCommand_BTCUSDT(t *testing.T) {
	req := httptest.NewRequest("GET", "/fetch-price/BTCUSDT", nil)
	rr := httptest.NewRecorder()
	handlers.HandleFetchPriceCommand(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}
}

// Test the status code 400 of FetchPrice endpoint as currency entered is invalid
func TestHandleFetchPriceCommand_BTCUSD(t *testing.T) {
	req := httptest.NewRequest("GET", "/fetch-price/BTCUSD", nil)
	rr := httptest.NewRecorder()
	handlers.HandleFetchPriceCommand(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadGateway, rr.Code)
	}
}
