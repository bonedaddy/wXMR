package service

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/bonedaddy/wxmr/db"
	"github.com/bonedaddy/wxmr/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/monero-ecosystem/go-monero-rpc-client/wallet"
	"github.com/stretchr/testify/require"
)

var (
	moneroWalletAddr = "http://127.0.0.1:6061/json_rpc"
	moneroWalletName = "monerot-estnet"
	serviceAddr      = "127.0.0.1:6666"
	gethRPC          = "http://127.0.0.1:8545"
	dbPath           = "test.db"
	ethKeyContents   = `{"address":"f2ea9ce3a27862650a8d40d98329dc0bd403a0c3","crypto":{"cipher":"aes-128-ctr","ciphertext":"75b1c0181fee4c7fa634a49bac74dacb12dcc27ec157e8549b6374924c5b1272","cipherparams":{"iv":"87117ca98fd0652db342a549889be7fe"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"032f744f960f549b4fc6d64177622fa1d34b16e8a4a106812d5ddc37487cb435"},"mac":"f468643296c390300e4ff808b9db4831275565c813c77cc64283adaaa1bcdffd"},"id":"c74bc71b-28a0-42cd-9474-4ffc98276d84","version":3}`
	ethAddress       = "0xf2ea9ce3a27862650a8d40d98329dc0bd403a0c3"
	id1              = "5d72a512cea06f80"
	addr1            = "A3jPvu6ZzrBNGUNyHTKU1G9LXcbMspjvtAQyTDdP4ADS4k3VMRudnrejc3w27gQWvPJbJEtKXkcde3yLM639RozHcBoehfZq8fQFVEwTmZ"
	tx1              = "5fa93c5c40e9841fe0982cecf921a1a187a89ca2e32a9d2feffb9912e5a5f174"
)

func TestService(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mc, err := rpc.NewClient(moneroWalletAddr)
	require.NoError(t, err)
	// ignore the error if it does not exist
	mc.CreateWallet(moneroWalletName)
	ec, err := ethclient.Dial(gethRPC)
	require.NoError(t, err)
	database, err := db.New(dbPath)
	require.NoError(t, err)
	t.Cleanup(func() {
		database.Close()
		ec.Close()
		mc.Close()
		os.Remove(dbPath)
	})
	auth, err := bind.NewTransactor(strings.NewReader(ethKeyContents), "")
	srv, err := NewService(serviceAddr, moneroWalletName, mc, ec, auth, database)
	require.NoError(t, err)

	go srv.Serve(ctx)

	// prepare the test deposit
	database.RegisterDeposit(addr1, id1)

	callURL := "http://" + serviceAddr

	hc := &http.Client{}
	req, err := http.NewRequest("GET", callURL+"/deposit_address", nil)
	require.NoError(t, err)

	resp, err := hc.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	var depResp ResponseGetDepositAddress
	require.NoError(t, json.Unmarshal(data, &depResp))
	t.Logf("deposit request: %s", depResp)
	/*!
	this is a temporary work around for a bug with the wallet rpc
	*/
	// revert this eventually once the bug is fixed
	depResp.Address = "BbcQw9uY7AeEP9eCqYL7D7GzJjyYD9DHPDc2dMHx4Wc3foC8UVcda6geg68jXYoqYo4Ku2KE4GqVm23fJkBmiP9uRfr4LyF"
	if resp, err := mc.Transfer(rpc.TransferOpts{
		Priority:     wallet.PriorityElevated,
		WalletName:   moneroWalletName,
		PaymentID:    depResp.PaymentID,
		Destinations: map[string]uint64{depResp.Address: wallet.Float64ToXMR(0.1)},
	}); err != nil {
		t.Fatal(err)
	} else {
		t.Log("transfer hash: ", resp.TxHash)
	}

	data, err = json.Marshal(&RequestMint{PaymentID: id1, EthAddress: ethAddress})
	require.NoError(t, err)

	req, err = http.NewRequest("POST", callURL+"/mint", bytes.NewReader(data))
	require.NoError(t, err)
	_, err = hc.Do(req)
	require.NoError(t, err)

	time.Sleep(time.Minute * 20)
}
