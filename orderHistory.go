package dovewallet

type OrderHistoryResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Result  []Order `json:"result,omitempty"`
}

type OrderHistoryParams struct {
	Status         string `url:"status,omitempty"`
	CurrencySymbol string `url:"currencySymbol,omitempty"`
	Market         string `url:"market"`
	WalletId       int64  `url:"walletId,omitempty"`
	Count          int    `url:"count,omitempty"`
	StartAt        jTime `url:"startAt,omitempty"`
}
