package testutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlockchain(t *testing.T) {
	blockchain := NewBlockchain(t)

	addr, xmr, err := blockchain.DeployReserve()
	require.NoError(t, err)

	t.Logf("contract addresss: %s\n", addr.String())

	require.NoError(t, blockchain.Mint(xmr, blockchain.auth.From, OneEthInWei))
}
