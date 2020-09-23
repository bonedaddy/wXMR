package service

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
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
	moneroWalletAddr       = "http://127.0.0.1:6061/json_rpc"
	moneroWalletName       = "monerot-estnet"
	serviceAddr            = "127.0.0.1:6666"
	gethRPC                = "http://127.0.0.1:8545"
	dbPath                 = "test.db"
	ethKeyContents         = `{"address":"f2ea9ce3a27862650a8d40d98329dc0bd403a0c3","crypto":{"cipher":"aes-128-ctr","ciphertext":"75b1c0181fee4c7fa634a49bac74dacb12dcc27ec157e8549b6374924c5b1272","cipherparams":{"iv":"87117ca98fd0652db342a549889be7fe"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"032f744f960f549b4fc6d64177622fa1d34b16e8a4a106812d5ddc37487cb435"},"mac":"f468643296c390300e4ff808b9db4831275565c813c77cc64283adaaa1bcdffd"},"id":"c74bc71b-28a0-42cd-9474-4ffc98276d84","version":3}`
	ethAddress             = "0xf2ea9ce3a27862650a8d40d98329dc0bd403a0c3"
	reserveContractAddress = "0x56D7C814700e6deD2Ab0173da2BdC57e1532bEd3"
	id1                    = "d1925a00c1625c3c"
	addr1                  = "A3jPvu6ZzrBNGUNyHTKU1G9LXcbMspjvtAQyTDdP4ADS4k3VMRudnrejc3w27gQWvPJbJEtKXkcde3yLM639RozHcGCtEPYWhbH7sbiFNN"
	tx1                    = "8226e4d3ab5a908e3fb91a8ad9def48444694764c953acf6826a8937467e7cf5"
)

func TestService(t *testing.T) {

	// set dev mode
	dev = true
	// overide dev confirmation count to 1
	devConfirmationCount = 1

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
	require.NoError(t, err)
	srv, err := NewService(ctx, serviceAddr, moneroWalletName, reserveContractAddress, mc, ec, auth, database)
	require.NoError(t, err)

	go srv.Serve(ctx)
	// start the hacky miner
	go ethMine(t, srv)
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
	we dont actually check this transaction due to the time the testnet takes to confirm blocks
	so we simply use the returned information to make sure we can send a transfer to it, all further
	testing occurs using a previously confirmed transaction
	*/
	if resp, err := mc.Transfer(rpc.TransferOpts{
		Priority:     wallet.PriorityElevated,
		WalletName:   moneroWalletName,
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
	resp, err = hc.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	log.Println("mint response: ", string(data))
	time.Sleep(time.Minute * 1)
}

// because we start the testenv geth node and set it
// to only mine blocks when there is a transaction, we need to
// create transactions to mine so that the tests can run properly
func ethMine(t *testing.T, s *Service) {
	for {
		select {
		case <-s.ctx.Done():
			return
		default:
		}
		time.Sleep(time.Second * 1)
		tx, err := s.rsrv.SetExchangeRate(s.auth, big.NewInt(1))
		if err != nil {
			t.Log("mine transaction send failed: ", err.Error())
			return
		}
		if _, err := bind.WaitMined(s.ctx, s.ec, tx); err != nil {
			t.Log("failed to wait for transaction to be mined: ", err)
			return
		}
	}
}
