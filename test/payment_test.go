package test

import (
	"hardiantojp/billing-service/routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMakePayment_Success(t *testing.T) {
	router := gin.Default()
	routes.RegisterRoutes(router)

	// create loan
	body := `{"loan_id": 201}`
	req, _ := http.NewRequest("POST", "/loan", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// make payment
	payReq, _ := http.NewRequest("POST", "/loan/201/pay", nil)
	payW := httptest.NewRecorder()
	router.ServeHTTP(payW, payReq)

	if payW.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", payW.Code)
	}
}

func TestMakePayment_LoanNotFound(t *testing.T) {
	router := gin.Default()
	routes.RegisterRoutes(router)

	payReq, _ := http.NewRequest("POST", "/loan/999/pay", nil)
	payW := httptest.NewRecorder()
	router.ServeHTTP(payW, payReq)

	if payW.Code != http.StatusNotFound {
		t.Errorf("expected 404 Not Found, got %d", payW.Code)
	}
}

func TestMakePayment_InvalidLoanID(t *testing.T) {
	router := gin.Default()
	routes.RegisterRoutes(router)

	payReq, _ := http.NewRequest("POST", "/loan/invalid/pay", nil)
	payW := httptest.NewRecorder()
	router.ServeHTTP(payW, payReq)

	if payW.Code != http.StatusBadRequest {
		t.Errorf("expected 400 Bad Request, got %d", payW.Code)
	}
}
