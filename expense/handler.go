package expense

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateExpense(s Services) gin.HandlerFunc {
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

func GetExpense(s Services) gin.HandlerFunc {
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

//func updateExpense() gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//	}
//}
//
//func getAllExpenses() gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//	}
//}
