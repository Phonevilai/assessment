package expense

type Services interface {
	Create(exp Req) (*Expense, error)
	//FindById(id int)
}

type MyService struct {
	*MyStore
}

func NewService(s *MyStore) *MyService {
	return &MyService{s}
}

func (s *MyService) Create(exp Req) (*Expense, error) {
	result, err := s.InsertExpense(exp)
	if err != nil {
		return nil, err
	}
	return result, nil
}
