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

type requestParams struct {
	Params []requestParam
}

type requestParam struct {
	Key string
	Value string
}

type DoveWalletClient interface {
	GetBalances() (balanceResponse BalancesResponse, err error)
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
func (d *DoveWallet) GetBalances() (balanceResponse BalancesResponse, err error) {
	r, err := d.client.do("GET", "account/getbalances", requestParams{}, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &balanceResponse)
	return
}

func (d *DoveWallet) GetOrderHistory(market string, walletId int64, count int, startAt *time.Time) (orderHistoryResponse OrderHistoryResponse, err error) {
	reqParams := requestParams{}

	if startAt != nil {
		reqParams.Params = append(reqParams.Params, requestParam{Key: "startat", Value:startAt.Format(TIME_FORMAT)})
	}

	reqParams.Params = append(reqParams.Params, requestParam{Key: "market", Value:market})

	resource := "account/getorderhistory"
	r, err := d.client.do("GET", resource, reqParams, true)
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
reqParams := requestParams{[]requestParam{{Key: "uuid", Value: uuid}}}

	resource := "account/getorder"

	r, err := d.client.do("GET", resource, reqParams, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(r, &orderResponse)
	return
}

func (d *DoveWallet) GetOpenOrders(market string, walletId int64) (orderHistoryResponse OrderHistoryResponse, err error) {
	resource := "market/getopenorders"

	reqParams := requestParams{[]requestParam{{Key: "market", Value: market}}}

	r, err := d.client.do("GET", resource, reqParams, true)
	if err != nil {
		return
	}
	fmt.Println(string(r))
	err = json.Unmarshal(r, &orderHistoryResponse)
	return
}

func (d *DoveWallet) OpenLimitOrder(direction, market string, quantity, rate decimal.Decimal, walletId int64) (limitOrderResponse OrderResponse, err error) {
	orderType := "buylimit"
	if direction == "SELL" {
		orderType = "selllimit"
	}

	resource :="market/" + orderType

	reqParams := requestParams{[]requestParam{
		{Key: "market", Value: market},
		{Key: "quantity", Value: quantity.String()},
		{Key: "rate", Value: rate.String()},
	},
	}

	r, err := d.client.do("GET", resource, reqParams, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(r, &limitOrderResponse)/**/
	return
}

func (rp requestParams) Len() int {
	return len(rp.Params)
}
func (rp requestParams) Swap(i, j int) {
	rp.Params[i], rp.Params[j] = rp.Params[j], rp.Params[i]
}
func (rp requestParams) Less(i, j int) bool {
	return rp.Params[i].Key < rp.Params[j].Key
}