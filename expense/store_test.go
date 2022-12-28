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
		e := ReqCreate{
			Title:  "strawberry smoothie",
			Amount: 1000,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"food", "beverage"},
		}
		got, _ := s.InsertExpense(e)
		fmt.Println("id", got.ID)
		assert.Equal(t, e.Title, got.Title)
	})

	t.Run("find expense bt id", func(t *testing.T) {
		getEnv()
		db := NewDB(os.Getenv("DATABASE_URL"))
		s := NewStore(db)
		result, _ := s.FindExpenseById("1")
		fmt.Println("result:", result)
		assert.Equal(t, 100, result.Amount)
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
		result, err := s.UpdateExpenseById(e)
		fmt.Println("result:", result)
		assert.Equal(t, nil, err)
	})

	t.Run("find all expenses", func(t *testing.T) {
		getEnv()
		db := NewDB(os.Getenv("DATABASE_URL"))
		s := NewStore(db)
		r, err := s.FindAllExpenses()
		fmt.Println("result:", len(r))
		assert.Equal(t, nil, err)
	})
}
