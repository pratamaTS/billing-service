package test

import (
	"hardiantojp/billing-service/routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestIsDelinquent_InitialFalse(t *testing.T) {
	router := gin.Default()
	routes.RegisterRoutes(router)

	// create loan
	body := `{"loan_id": 401}`
	req, _ := http.NewRequest("POST", "/loan", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// check delinquent
	getReq, _ := http.NewRequest("GET", "/loan/401/is_delinquent", nil)
	getW := httptest.NewRecorder()
	router.ServeHTTP(getW, getReq)

	if getW.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", getW.Code)
	}
}

func TestIsDelinquent_LoanNotFound(t *testing.T) {
	router := gin.Default()
	routes.RegisterRoutes(router)

	getReq, _ := http.NewRequest("GET", "/loan/999/is_delinquent", nil)
	getW := httptest.NewRecorder()
	router.ServeHTTP(getW, getReq)

	if getW.Code != http.StatusNotFound {
		t.Errorf("expected 404 Not Found, got %d", getW.Code)
	}
}
