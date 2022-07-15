package models

type User struct {
	ID            int    `json:"user_id,omitempty"`
	Name          string `json:"name,omitempty"`
	Username      string `json:"username,omitempty"`
	Type          string `json:"type,omitempty"`
	SavingBalance int    `json:"saving_balance,omitempty"`
	LoanBalance   int    `json:"laoan_balance,omitempty"`
}

type UserSaving struct {
	User   User   `json:"user"`
	Saving Saving `json:"saving data"`
}
