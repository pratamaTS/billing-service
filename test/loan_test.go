package test

import (
	"hardiantojp/billing-service/routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateLoan_Success(t *testing.T) {
	router := gin.Default()
	routes.RegisterRoutes(router)

	body := `{"loan_id": 101}`
	req, _ := http.NewRequest("POST", "/loan", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", w.Code)
	}
}

func TestCreateLoan_InvalidPayload(t *testing.T) {
	router := gin.Default()
	routes.RegisterRoutes(router)
	w := httptest.NewRecorder()
	reqBody := `{"loan_id": "invalid"}`

	req, _ := http.NewRequest("POST", "/loan", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 Bad Request, got %d", w.Code)
	}
}
