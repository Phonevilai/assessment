package store

import (
	"database/sql"
	"github.com/Phonevilai/assessment/pkg/expense"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

type DataStore interface {
	CreateExpense(e expense.Expense) error
	FindExpenseById(findId int) (*expense.Expense, error)
	UpdateExpense(e expense.Expense) (*expense.Expense, error)
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

func (s *Store) UpdateExpenseByID(e expense.Expense) (*expense.Expense, error) {
	row := s.db.QueryRow("UPDATE expenses SET title = $1, amount = $2, note = $3, tags = $4 WHERE id = $5 RETURNING id, title, amount, note, tags",
		e.Title, e.Amount, e.Note, pq.Array(&e.Tags), e.ID)
	var id int
	var ne expense.Expense
	if err := row.Scan(&id, &ne.Title, &ne.Amount, &ne.Note, (*pq.StringArray)(&ne.Tags)); err != nil {
		return nil, err
	}
	ne.ID = strconv.Itoa(id)

	return &ne, nil
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
