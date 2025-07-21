package models

import "time"

type Payment struct {
	Week   int       `json:"week"`
	Amount int       `json:"amount"`
	Paid   bool      `json:"paid"`
	PaidAt time.Time `json:"paid_at,omitempty"`
}

type Loan struct {
	LoanID       int       `json:"loan_id"`
	TotalAmount  int       `json:"total_amount"`
	Weeks        int       `json:"weeks"`
	InterestRate float64   `json:"interest_rate"`
	Schedule     []Payment `json:"schedule"`
}
