package expense

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Expense struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

type Router struct {
	*gin.Engine
}

func NewHandler() *Router {
	r := gin.Default()
	r.GET("/healthz", healthCheck())
	return &Router{r}
}

func healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
