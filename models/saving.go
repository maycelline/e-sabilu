package models

import "time"

type Saving struct {
	UserID   int       `json:"user_id,omitempty"`
	Date     time.Time `json:"date,omitempty"`
	Sw       int       `json:"simpanan_wajib,omitempty"`
	Ss       int       `json:"simpanan_s?,omitempty"`
	Withdraw int       `json:"ditarik,omitempty"`
	// Balance  int       `json:"saldo,omitempty"`
	Jsp int `json:"jumlah_saham_penuh,omitempty"`
}
