package expense

import (
	"github.com/gin-gonic/gin"
)

func NewMainHandler() *Router {
	r := gin.Default()
	r.GET("/healthz", healthCheck())

	r.POST("/expenses", createExpense())

	return &Router{r}
}

func createExpense() gin.HandlerFunc {

}

func getExpense() gin.HandlerFunc {

}

func updateExpense() gin.HandlerFunc {

}

func getAllExpenses() {

}