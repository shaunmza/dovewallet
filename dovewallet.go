package dovewallet

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/shopspring/decimal"
	"time"
)

const (
	API_BASE    = "https://api.dovewallet.com/"
	API_VERSION = "v1.1"
	TIME_FORMAT = "2006-01-02T15:04:05"
)

// New returns an instantiated dovewallet struct
func New(apiKey, apiSecret string) *DoveWallet {
	client := NewClient(apiKey, apiSecret)
	return &DoveWallet{client}
}

// dovewallet represents a Dove Wallet client
type DoveWallet struct {
	client *client
}

type DoveWalletClient interface {
	GetBalances() (balanceResponse BalanceResponse, err error)
	GetOrderHistory(market string, walletId int64, count int, startAt *time.Time) (orderHistoryResponse OrderHistoryResponse, err error)
	GetOrder(uuid string) (orderResponse OrderResponse, err error)
	GetOpenOrders(market string, walletId int64) (orderHistoryResponse OrderHistoryResponse, err error)
	OpenLimitOrder(direction, market string, quantity, rate decimal.Decimal, walletId int64) (limitOrderResponse OrderResponse, err error)
}

// set enable/disable http request/response dump
func (c *DoveWallet) SetDebug(enable bool) {
	c.client.debug = enable
}

// Account

// GetBalances is used to retrieve all balances from your account
func (d *DoveWallet) GetBalances() (balanceResponse BalanceResponse, err error) {
	r, err := d.client.do("GET", "account/getbalances", "", true)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &balanceResponse)
	return
}

func (d *DoveWallet) GetOrderHistory(market string, walletId int64, count int, startAt *time.Time) (orderHistoryResponse OrderHistoryResponse, err error) {
	from := ""
	if startAt != nil {
		from = startAt.Format(TIME_FORMAT)
	}

	var params = OrderHistoryParams{
		Market:   market,
		WalletId: walletId,
		Count:    count,
		StartAt:  from,
	}
	v, _ := query.Values(params)
	queryParams := v.Encode()
	resource := "account/getorderhistory"
	if len(queryParams) != 0 {
		resource += "?"
	}
	r, err := d.client.do("GET", resource+queryParams, "", true)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &orderHistoryResponse)
	return
}

func (d *DoveWallet) GetOrder(uuid string) (orderResponse OrderResponse, err error) {

	resource := "account/getorder?uuid=" + uuid

	r, err := d.client.do("GET", resource, "", true)
	if err != nil {
		return
	}

	err = json.Unmarshal(r, &orderResponse)
	return
}

func (d *DoveWallet) GetOpenOrders(market string, walletId int64) (orderHistoryResponse OrderHistoryResponse, err error) {

	var params = OpenOrderParams{
		Market:   market,
		WalletId: walletId,
	}
	v, _ := query.Values(params)
	queryParams := v.Encode()
	resource := "market/getopenorders"
	if len(queryParams) != 0 {
		resource += "?"
	}
	r, err := d.client.do("GET", resource+queryParams, "", true)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &orderHistoryResponse)
	return
}

func (d *DoveWallet) OpenLimitOrder(direction, market string, quantity, rate decimal.Decimal, walletId int64) (limitOrderResponse OrderResponse, err error) {

	orderType := "buylimit"
	if direction == "SELL" {
		orderType = "selllimit"
	}

	r, err := d.client.do("GET", fmt.Sprintf("market/%s?market=%s&quantity=%s&rate=%s", orderType, market, quantity.String(), rate.String()), "", true)
	if err != nil {
		return
	}

	err = json.Unmarshal(r, &limitOrderResponse)
	return
}
