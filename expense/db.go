package expense

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
)

type Storer interface {
	InsertExpense(exp Req) (*Expense, error)
	FindExpenseById(findId int) (*Expense, error)
	UpdateExpenseById(e Expense) (*Expense, error)
	FindAllExpenses() ([]*Expense, error)
}

type MyStore struct {
	*sql.DB
}

func NewStore(db *sql.DB) *MyStore {
	return &MyStore{db}
}

func (s *MyStore) InsertExpense(exp Req) (*Expense, error) {
	insert := `
INSERT INTO expenses (title, amount, note, tags) VALUES ($1, $2, $3, $4) RETURNING expenses.id, title, amount, note, tags;
`
	var ne Expense
	err := s.QueryRow(insert, exp.Title, exp.Amount, exp.Note, pq.Array(&exp.Tags)).Scan(&ne.ID, &ne.Title, &ne.Amount, &ne.Note, (*pq.StringArray)(&ne.Tags))
	if err != nil {
		return nil, err
	}

	return &ne, nil
}

func (s *MyStore) FindExpenseById(id string) (*Expense, error) {
	rows, err := s.Query("SELECT * FROM expenses WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var e Expense
	for rows.Next() {
		err = rows.Scan(&e.ID, &e.Title, &e.Amount, &e.Note, (*pq.StringArray)(&e.Tags))
		if err != nil {
			return nil, err
		}
	}

	return &e, nil
}

func (s *MyStore) UpdateExpenseById(e Expense) (*Expense, error) {
	update := `
UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE expenses.id = $1 RETURNING expenses.id, title, amount, note, tags;
`
	var ne Expense
	err := s.QueryRow(update, e.ID, e.Title, e.Amount, e.Note, pq.Array(&e.Tags)).Scan(&ne.ID, &ne.Title, &ne.Amount, &ne.Note, (*pq.StringArray)(&ne.Tags))
	if err != nil {
		return nil, err
	}

	fmt.Println("Updated:", ne.ID, ne.Title, ne.Amount, ne.Note, ne.Tags)
	return &ne, nil
}

func (s *MyStore) FindAllExpenses() ([]*Expense, error) {
	stml, err := s.Prepare("SELECT * FROM expenses ORDER BY expenses.id ASC")
	if err != nil {
		return nil, err
	}
	defer stml.Close()

	rows, err := stml.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []*Expense
	for rows.Next() {
		var e Expense
		err = rows.Scan(&e.ID, &e.Title, &e.Amount, &e.Note, (*pq.StringArray)(&e.Tags))
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, &e)
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
