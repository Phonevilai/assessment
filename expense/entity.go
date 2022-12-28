package expense

type Req struct {
	Title  string   `json:"title" binding:"required"`
	Amount int      `json:"amount" binding:"required"`
	Note   string   `json:"note" binding:"required"`
	Tags   []string `json:"tags" binding:"required"`
}

type Res struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Amount int      `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}
