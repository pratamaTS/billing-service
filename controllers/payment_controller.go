package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func MakePayment(c *gin.Context) {
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

	for i := 0; i < len(loan.Schedule); i++ {
		if !loan.Schedule[i].Paid {
			loan.Schedule[i].Paid = true
			loan.Schedule[i].PaidAt = time.Now()
			c.JSON(http.StatusOK, gin.H{"message": "payment successful"})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "loan already paid"})
}
