package store

import (
	"database/sql"
	"fmt"
	"github.com/Phonevilai/assessment/pkg/expense"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
	"strconv"
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

func (s *Store) FindExpenseById(findId int) (*expense.Expense, error) {
	rows, err := s.db.Query("SELECT * FROM expenses WHERE id = $1", findId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var e expense.Expense
	var id int
	for rows.Next() {
		err = rows.Scan(&id, &e.Title, &e.Amount, &e.Note, (*pq.StringArray)(&e.Tags))
		if err != nil {
			return nil, err
		}
		e.ID = strconv.Itoa(id)
	}
	return &e, nil
}

func (s *Store) UpdateExpenseById(e expense.Expense) (*expense.Expense, error) {
	stmt, err := s.db.Prepare("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE expenses.id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	inId, err := strconv.Atoi(e.ID)
	if err != nil {
		return nil, err
	}
	if _, err = stmt.Exec(inId, e.Title, e.Amount, e.Note, pq.Array(&e.Tags)); err != nil {
		return nil, err
	}
	fmt.Println("update success")
	return &e, nil
}

func (s *Store) FindAllExpenses() ([]*expense.Expense, error) {
	stml, err := s.db.Prepare("SELECT * FROM expecses")
	if err != nil {
		return nil, err
	}
	rows, err := stml.Query()
	if err != nil {
		return nil, err
	}

	var expenses []*expense.Expense
	for rows.Next() {
		var e *expense.Expense
		var id int
		err = rows.Scan(&id, &e.Title, &e.Amount, &e.Note, (*pq.StringArray)(&e.Tags))
		if err != nil {
			return nil, err
		}
		e.ID = strconv.Itoa(id)
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
