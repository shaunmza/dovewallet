package dovewallet

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	API_KEY    = ""
	API_SECRET = ""
)

func TestClient_GetBalances(t *testing.T) {
	r := NewClient(API_KEY, API_SECRET)

	c := DoveWallet{r}

	balanceResponse, err := c.GetBalances()
	require.NoError(t, err)
	require.Equal(t, true, balanceResponse.Success)
}

func TestClient_GetOpenOrders(t *testing.T) {
	r := NewClient(API_KEY, API_SECRET)

	c := DoveWallet{r}

	orderResponse, err := c.GetOpenOrders("USDT-BTC", 0)
	require.NoError(t, err)
	require.Equal(t, true, orderResponse.Success)
}

func TestClient_GetOrderHistory(t *testing.T) {
	r := NewClient(API_KEY, API_SECRET)

	c := DoveWallet{r}

	orderResponse, err := c.GetOrderHistory("USDT-BTC", 0, 0, nil)
	require.NoError(t, err)
	require.Equal(t, true, orderResponse.Success)
}
