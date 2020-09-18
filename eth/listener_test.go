package eth

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/bonedaddy/wxmr/bindings/reserve"
	"github.com/bonedaddy/wxmr/testutils"
	"github.com/bonedaddy/wxmr/utils"
	"github.com/stretchr/testify/require"
)

func TestListener(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	blockchain := testutils.NewBlockchain(t)

	addr, contract, err := blockchain.DeployReserve()
	require.NoError(t, err)

	t.Log("contract address: ", addr.String())
	ch := make(chan *reserve.ReserveCoinsBurned)
	go func() {
		listener := NewListener(contract)
		listener.HandleCoinBurn(ctx, ch)
	}()
	go func() {
		for i := 0; i < 100; i++ {
			require.NoError(t, blockchain.Mint(contract, blockchain.Auth().From, testutils.OneEthInWei))
		}

		for i := 0; i < 100; i++ {
			hash := utils.SumKeccak256([]byte(fmt.Sprintf("%v", i)))
			require.NoError(t, blockchain.Burn(contract, testutils.OneEthInWei, hash))
		}
	}()
	for {
		select {
		case event := <-ch:
			log.Print("new event: ", event)
		}
	}
}
