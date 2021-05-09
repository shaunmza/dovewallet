package dovewallet

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

const (
	API_BASE    = "https://api.dovewallet.com/"
	API_VERSION = "v1.1"
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
		from = "&startat=" + startAt.Format(TIME_FORMAT)
	}

	resource := "account/getorderhistory"
	r, err := d.client.do("GET", resource+fmt.Sprintf("?market=%s%s", market, from), "", true)
	if err != nil {
		return
	}
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
	resource := "market/getopenorders"

	r, err := d.client.do("GET", resource+fmt.Sprintf("?market=%s", market), "", true)
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
