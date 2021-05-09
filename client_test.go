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
	require.Equal(t, 3, len(balanceResponse.Result))
}
