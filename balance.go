package dovewallet

import (
	"github.com/shopspring/decimal"
)

type BalanceResponse struct {
	Currency bool      `json:"success"`
	Message  string    `json:"message"`
	Result   []Balance `json:"result"`
}

type Balance struct {
	Currency      string          `json:"Currency"`
	Balance       decimal.Decimal `json:"Balance"`
	Available     decimal.Decimal `json:"Available"`
	Pending       decimal.Decimal `json:"Pending"`
	CryptoAddress string          `json:"CryptoAddress"`
	Requested     bool            `json:"Requested"`
	Uuid          string          `json:"Uuid"`
}
