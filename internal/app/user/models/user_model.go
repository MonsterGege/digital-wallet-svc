package models

type UserParam struct {
	ID               int     `form:"id" json:"id" binding:"required"`
	Name             string  `form:"name" json:"name"`
	WalletID         int     `form:"wallet_id" json:"wallet_id"`
	WithdrawalAmount float64 `form:"withdrawal_amount" json:"withdrawal_amount"`
}

type ResponseBalance struct {
	Balance float64 `json:"balance"`
}
