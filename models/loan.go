package models

import "time"

type Loan struct {
	UserID      int       `json:"user_id,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	Amount      int       `json:"jumlah_pinjaman,omitempty"`
	Installment int       `json:"angsuran,omitempty"`
	// Balance     int       `json:"saldo,omitempty"`
	Interest int `json:"bunga,omitempty"`
	Fine     int `json:"denda,omitempty"`
}
