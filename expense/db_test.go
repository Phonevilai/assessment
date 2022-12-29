package expense

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func getEnv() {
	if err := godotenv.Load("../dev.env"); err != nil {
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
		e := Req{
			Title:  "strawberry smoothie",
			Amount: 1000,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"food", "beverage"},
		}
		got, err := s.InsertExpense(e)
		if assert.NoError(t, err) {
			assert.Equal(t, e.Title, got.Title)
		}
	})

	t.Run("find expense bt id", func(t *testing.T) {
		getEnv()
		db := NewDB(os.Getenv("DATABASE_URL"))
		s := NewStore(db)
		got, err := s.FindExpenseById("1")

		// Assertions
		if assert.NoError(t, err) {
			assert.Equal(t, "1", got.ID)
		}
	})

	t.Run("update expense by id and returning", func(t *testing.T) {
		getEnv()
		db := NewDB(os.Getenv("DATABASE_URL"))
		s := NewStore(db)
		e := Expense{
			ID:     "1",
			Title:  "toooky",
			Amount: 100.00,
			Note:   "1",
			Tags:   []string{"123"},
		}

		got, err := s.UpdateExpenseById(e)
		if assert.NoError(t, err) {
			assert.Equal(t, "1", got.ID)
		}
	})

	t.Run("find all expenses", func(t *testing.T) {
		getEnv()
		db := NewDB(os.Getenv("DATABASE_URL"))
		s := NewStore(db)
		r, err := s.FindAllExpenses()

		if assert.NoError(t, err) {
			if assert.NotEqual(t, 0, len(r)) {
				assert.Condition(t, func() bool {
					if len(r) <= 0 {
						return false
					}
					return true
				})
			}
		}
	})
}
