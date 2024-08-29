package types

type AppSubscription struct {
	ID        string `json:"id"`
	ReturnURL string `json:"returnUrl"`
	Status    string `json:"status"`
}

type DiscountCodeNode struct {
	ID string `json:"id"`
}

type UserError struct {
	Code      string   `json:"code"`
	ExtraInfo string   `json:"extraInfo"`
	Field     []string `json:"field"`
	Message   string   `json:"message"`
}
