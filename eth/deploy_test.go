package eth

import (
	"context"
	"strings"
	"testing"

	"github.com/bonedaddy/wxmr/bindings/reserve"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

var (
	gethRPC        = "http://127.0.0.1:8545"
	ethKeyContents = `{"address":"f2ea9ce3a27862650a8d40d98329dc0bd403a0c3","crypto":{"cipher":"aes-128-ctr","ciphertext":"75b1c0181fee4c7fa634a49bac74dacb12dcc27ec157e8549b6374924c5b1272","cipherparams":{"iv":"87117ca98fd0652db342a549889be7fe"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"032f744f960f549b4fc6d64177622fa1d34b16e8a4a106812d5ddc37487cb435"},"mac":"f468643296c390300e4ff808b9db4831275565c813c77cc64283adaaa1bcdffd"},"id":"c74bc71b-28a0-42cd-9474-4ffc98276d84","version":3}`
)

func TestDeployReserveContract(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	auth, err := bind.NewTransactor(strings.NewReader(ethKeyContents), "")
	require.NoError(t, err)
	ec, err := ethclient.Dial(gethRPC)
	require.NoError(t, err)
	defer ec.Close()
	addr, contract, err := DeployReserveContract(ctx, auth, ec)
	require.NoError(t, err)
	admin, err := contract.Admin(nil)
	require.NoError(t, err)
	require.Equal(t, auth.From, admin)
	contract2, err := reserve.NewReserve(addr, ec)
	require.NoError(t, err)
	admin2, err := contract2.Admin(nil)
	require.NoError(t, err)
	require.Equal(t, admin, admin2)
}
