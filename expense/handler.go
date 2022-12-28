package expense

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(s *MyService) *gin.Engine {

	r := gin.Default()
	r.GET("/healthz", healthCheck())
	r.POST("/expenses", createExpense(s))
	r.GET("/expenses/:id", getExpense(s))
	r.PUT("/expenses/:id", updateExpense(s))
	r.GET("/expenses", getAllExpenses(s))

	return r
}

func createExpense(s Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Req
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		r, err := s.Create(req)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusCreated, r)
	}
}

func getExpense(s Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		r, err := s.FindById(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		if r.ID == "" {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, r)
	}
}

func updateExpense(s Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var req Req
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		r, err := s.Update(id, req)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, r)
	}
}

func getAllExpenses(s Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, err := s.GetAll()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, r)
	}
}

func healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
