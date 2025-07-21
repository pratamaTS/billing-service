package routes

import (
	"hardiantojp/billing-service/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/loan", controllers.CreateLoan)
	r.GET("/loan/:id/outstanding", controllers.GetOutstanding)
	r.GET("/loan/:id/is_delinquent", controllers.IsDelinquent)
	r.POST("/loan/:id/pay", controllers.MakePayment)
}
