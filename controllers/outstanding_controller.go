package controllers

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOutstanding(c *gin.Context) {
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

	out := 0
	for _, p := range loan.Schedule {
		if !p.Paid {
			out += p.Amount
		}
	}
	c.JSON(http.StatusOK, gin.H{"outstanding": out})
}
