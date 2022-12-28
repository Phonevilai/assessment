package expense

type Expense struct {
	ID     string
	Title  string
	Amount int
	Note   string
	Tags   []string
}
