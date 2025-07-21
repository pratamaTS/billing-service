package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IsDelinquent(c *gin.Context) {
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

	missed := 0
	for i := len(loan.Schedule) - 1; i >= 0; i-- {
		if loan.Schedule[i].Paid {
			break
		}
		missed++
	}
	c.JSON(http.StatusOK, gin.H{"is_delinquent": missed >= 2})
}
