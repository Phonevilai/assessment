package expense

type Services interface {
	Create(exp Req) (*Res, error)
	FindById(id string) (*Res, error)
	Update(id string, exp Req) (*Res, error)
}

type MyService struct {
	*MyStore
}

func NewService(s *MyStore) *MyService {
	return &MyService{s}
}

func (s *MyService) Create(exp Req) (*Res, error) {
	r, err := s.InsertExpense(exp)
	if err != nil {
		return nil, err
	}
	return &Res{
		ID:     r.ID,
		Title:  r.Title,
		Amount: r.Amount,
		Note:   r.Note,
		Tags:   r.Tags,
	}, nil
}

func (s *MyService) FindById(id string) (*Res, error) {
	r, err := s.FindExpenseById(id)
	if err != nil {
		return nil, err
	}
	return &Res{
		ID:     r.ID,
		Title:  r.Title,
		Amount: r.Amount,
		Note:   r.Note,
		Tags:   r.Tags,
	}, nil
}

func (s *MyService) Update(id string, exp Req) (*Res, error) {
	r, err := s.UpdateExpenseById(Expense{
		ID:     id,
		Title:  exp.Title,
		Amount: exp.Amount,
		Note:   exp.Note,
		Tags:   exp.Tags,
	})
	if err != nil {
		return nil, err
	}
	return &Res{
		ID:     r.ID,
		Title:  r.Title,
		Amount: r.Amount,
		Note:   r.Note,
		Tags:   r.Tags,
	}, nil
}
