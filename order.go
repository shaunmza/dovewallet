package dovewallet

import (
	"github.com/shopspring/decimal"
)

type OrderResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  Order  `json:"result"`
}

type OpenOrderParams struct {
	Market   string `json:"market"`
	WalletId int64  `json:"walletid"`
}

type LimitOrderParams struct {
	Market   string          `json:"market"`
	Quantity decimal.Decimal `json:"quantity"`
	Rate     decimal.Decimal `json:"rate"`
	WalletId int64           `json:"walletid"`
	Magic    int             `json:"magic"`
}

type LimitOrderResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Result  OrderId `json:"result"`
}
type OrderId struct {
	Uuid int64 `json:"uuid"`
}

type Order struct {
	AccountId         string          `json:"AccountId"`
	Uuid              string          `json:"Uuid"`
	OrderUuid         string          `json:"OrderUuid"`
	Exchange          string          `json:"Exchange"`
	TimeStamp         jTime           `json:"TimeStamp,omitempty"`
	OrderType         string          `json:"OrderType"`
	Limit             decimal.Decimal `json:"Limit"`
	Magic             int             `json:"Magic"`
	Quantity          decimal.Decimal `json:"Quantity"`
	QuantityRemaining decimal.Decimal `json:"QuantityRemaining"`
	Commission        decimal.Decimal `json:"Commission"`
	Price             decimal.Decimal `json:"Price"`
	PricePerUnit      decimal.Decimal `json:"PricePerUnit"`
	IsConditional     bool            `json:"IsConditional"`
	Opened            jTime           `json:"Opened,omitempty"`
	Closed            *jTime          `json:"Closed,omitempty"`
	IsOpen            bool            `json:"IsOpen"`
}
