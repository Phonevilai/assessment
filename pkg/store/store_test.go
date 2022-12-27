package store

import (
	"fmt"
	"github.com/Phonevilai/assessment/pkg/expense"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func getEnv() {
	if err := godotenv.Load("../../dev.env"); err != nil {
		fmt.Printf("please consider environment variables: %s\n", err)
	}
}

func TestConnectToDB(t *testing.T) {

	t.Run("create table", func(t *testing.T) {
		getEnv()
		got := NewDB(os.Getenv("DATABASE_URL"))
		assert.NotNil(t, got)
	})

	t.Run("create expense", func(t *testing.T) {
		getEnv()
		db := NewDB(os.Getenv("DATABASE_URL"))
		s := NewStore(db)
		e := expense.Expense{
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"food", "beverage"},
		}
		got := s.CreateExpense(e)
		assert.Equal(t, nil, got)
	})

	t.Run("find expense bt id", func(t *testing.T) {
		getEnv()
		db := NewDB(os.Getenv("DATABASE_URL"))
		s := NewStore(db)
		result, _ := s.FindExpenseById(1)
		fmt.Println("result:", result)
		assert.Equal(t, 79.00, result.Amount)
	})

	t.Run("update expense by id", func(t *testing.T) {
		getEnv()
		db := NewDB(os.Getenv("DATABASE_URL"))
		s := NewStore(db)
		e := expense.Expense{
			ID:     "4",
			Title:  "strawberry smoothie",
			Amount: 5.00,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"beverage"},
		}
		result, err := s.UpdateExpenseById(e)
		fmt.Println("result:", result)
		assert.Equal(t, nil, err)
	})
}
