package controllers

import (
	"hardiantojp/billing-service/models"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	loans = map[int]*models.Loan{}
	mu    sync.Mutex
)

func CreateLoan(c *gin.Context) {
	type Req struct {
		LoanID int `json:"loan_id"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	loanAmount := 5_000_000
	weeks := 50
	rate := 0.10
	totalPayable := int(float64(loanAmount) * (1 + rate))
	weekly := totalPayable / weeks

	schedule := make([]models.Payment, weeks)
	for i := 0; i < weeks; i++ {
		schedule[i] = models.Payment{
			Week:   i + 1,
			Amount: weekly,
			Paid:   false,
		}
	}

	loan := &models.Loan{
		LoanID:       req.LoanID,
		TotalAmount:  loanAmount,
		Weeks:        weeks,
		InterestRate: rate,
		Schedule:     schedule,
	}
	loans[req.LoanID] = loan

	c.JSON(http.StatusOK, loan)
}

func GetLoan(c *gin.Context) {
	loanID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan id"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	loan, ok := loans[loanID]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "loan not found"})
		return
	}

	c.JSON(http.StatusOK, loan)
}
