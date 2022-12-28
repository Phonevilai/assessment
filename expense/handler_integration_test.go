//go:build integration

package expense

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestCreateExpense(t *testing.T) {
	//init env
	getEnv()

	// Arrange
	db := NewDB(os.Getenv("DATABASE_URL"))
	mydb := NewStore(db)
	service := NewService(mydb)
	r := gin.Default()
	r.POST("/expenses", createExpense(service))

	// Act
	expected := http.StatusCreated
	input := strings.NewReader(`{
    "title": "strawberry smoothie",
    "amount": 79,
    "note": "night market promotion discount 10 bath", 
    "tags": ["food", "beverage"]
}`)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/expenses", input)
	r.ServeHTTP(w, req)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, expected, w.Code)
	}
}

func TestGetExpenseById(t *testing.T) {
	//init env
	getEnv()

	// Arrange
	db := NewDB(os.Getenv("DATABASE_URL"))
	mydb := NewStore(db)
	service := NewService(mydb)
	r := gin.Default()
	r.GET("/expenses/:id", getExpense(service))

	// Act
	expected := http.StatusOK
	input := "1"
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/expenses/%s", input), nil)
	r.ServeHTTP(w, req)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, expected, w.Code)
	}
}

func TestUpdateExpenseById(t *testing.T) {
	//init env
	getEnv()

	// Arrange
	db := NewDB(os.Getenv("DATABASE_URL"))
	mydb := NewStore(db)
	service := NewService(mydb)
	r := gin.Default()
	r.PUT("/expenses/:id", updateExpense(service))

	// Act
	expected := http.StatusOK
	id := "1"
	body := strings.NewReader(`{
    "title": "Orange",
    "amount": 100,
    "note": "no discount", 
    "tags": ["fruit"]
}`)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", fmt.Sprintf("/expenses/%s", id), body)
	r.ServeHTTP(w, req)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, expected, w.Code)
	}
}

func TestGetAllExpenses(t *testing.T) {
	//init env
	getEnv()

	// Arrange
	db := NewDB(os.Getenv("DATABASE_URL"))
	mydb := NewStore(db)
	service := NewService(mydb)
	r := gin.Default()
	r.GET("/expenses", getAllExpenses(service))

	// Act
	expected := http.StatusOK
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/expenses", nil)
	r.ServeHTTP(w, req)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, expected, w.Code)
	}
}
