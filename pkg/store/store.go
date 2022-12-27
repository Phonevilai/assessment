package store

import (
	"database/sql"
	"github.com/Phonevilai/assessment/pkg/expense"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateExpense(e expense.Expense) error {
	row := s.db.QueryRow("INSERT INTO expenses (title, amount, note, tags) VALUES ($1, $2, $3, $4) RETURNING id", e.Title, e.Amount, e.Note, pq.Array(&e.Tags))
	var id int
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (s *Store) FindExpenseById(id string) (expense.Expense, error) {

}

func (s *Store) FindAllExpenses() ([]expense.Expense, error) {
	stml, err := s.db.Prepare("SELECT * FROM expecses")
	if err != nil {
		return nil, err
	}
	rows, err := stml.Query()
	if err != nil {
		return nil, err
	}

	var expenses []expense.Expense
	for rows.Next() {
		var e expense.Expense
		var id int
		err = rows.Scan(&id, &e.Title, &e.Amount, &e.Note, &e.Tags)
		if err != nil {
			return nil, err
		}
		e.ID = strconv.Itoa(id)
		expenses = append(expenses, e)
	}

	return expenses, nil
}

func NewDb(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error ", err)
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table ", err)
	}

	return db
}
