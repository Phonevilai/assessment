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
		c.JSON(http.StatusCreated, Res{
			ID:     r.ID,
			Title:  r.Title,
			Amount: r.Amount,
			Note:   r.Note,
			Tags:   r.Tags,
		})
	}
}

//
//func GetExpense(s Services) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		id := c.Param("id")
//
//	}
//}

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
