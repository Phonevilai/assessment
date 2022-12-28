package store

import (
	"database/sql"
	"fmt"
	"github.com/Phonevilai/assessment/pkg/expense"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
)

type DataStore interface {
	CreateExpense(e expense.Expense) error
	FindExpenseById(findId int) (*expense.Expense, error)
	UpdateExpenseById(e expense.Expense) (*expense.Expense, error)
	FindAllExpenses() ([]*expense.Expense, error)
}

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

func (s *Store) FindExpenseById(id string) (*expense.Expense, error) {
	rows, err := s.db.Query("SELECT * FROM expenses WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var e expense.Expense

	for rows.Next() {
		err = rows.Scan(&e.ID, &e.Title, &e.Amount, &e.Note, (*pq.StringArray)(&e.Tags))
		if err != nil {
			return nil, err
		}
	}

	return &e, nil
}

func (s *Store) UpdateExpenseById(e expense.Expense) (*expense.Expense, error) {
	update := `
UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE expenses.id = $1 RETURNING expenses.id, title, amount, note, tags;
`
	var ne expense.Expense
	err := s.db.QueryRow(update, e.ID, e.Title, e.Amount, e.Note, pq.Array(&e.Tags)).Scan(&ne.ID, &ne.Title, &ne.Amount, &ne.Note, (*pq.StringArray)(&ne.Tags))
	if err != nil {
		return nil, err
	}

	fmt.Println("Updated:", ne.ID, ne.Title, ne.Amount, ne.Note, ne.Tags)
	return &ne, nil
}

func (s *Store) FindAllExpenses() ([]expense.Expense, error) {
	stml, err := s.db.Prepare("SELECT * FROM expenses")
	if err != nil {
		return nil, err
	}
	defer stml.Close()

	rows, err := stml.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var expenses []expense.Expense
	for rows.Next() {
		var e expense.Expense
		err = rows.Scan(&e.ID, &e.Title, &e.Amount, &e.Note, (*pq.StringArray)(&e.Tags))
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, e)
	}

	return expenses, nil
}

func NewDB(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error ", err)
	}

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
