package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/bonedaddy/wxmr/db"
	"github.com/bonedaddy/wxmr/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

var (
	moneroWalletAddr = "http://127.0.0.1:6061/json_rpc"
	moneroWalletName = "monerot-estnet"
	serviceAddr      = "127.0.0.1:6666"
	gethRPC          = "http://127.0.0.1:8545"
	dbPath           = "test.db"
	ethKeyContents   = `{"address":"f2ea9ce3a27862650a8d40d98329dc0bd403a0c3","crypto":{"cipher":"aes-128-ctr","ciphertext":"75b1c0181fee4c7fa634a49bac74dacb12dcc27ec157e8549b6374924c5b1272","cipherparams":{"iv":"87117ca98fd0652db342a549889be7fe"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"032f744f960f549b4fc6d64177622fa1d34b16e8a4a106812d5ddc37487cb435"},"mac":"f468643296c390300e4ff808b9db4831275565c813c77cc64283adaaa1bcdffd"},"id":"c74bc71b-28a0-42cd-9474-4ffc98276d84","version":3}`
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
	})
	auth, err := bind.NewTransactor(strings.NewReader(ethKeyContents), "")
	srv, err := NewService(serviceAddr, moneroWalletName, mc, ec, auth, database)
	require.NoError(t, err)

	go srv.Serve(ctx)
	time.Sleep(time.Second)

	callURL := "http://" + serviceAddr
	hc := &http.Client{}
	req, err := http.NewRequest("GET", callURL+"/deposit_address", nil)
	require.NoError(t, err)
	resp, err := hc.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	log.Println(string(data))
	var depResp ResponseGetDepositAddress
	require.NoError(t, json.Unmarshal(data, &depResp))

}
