package service

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/bonedaddy/wxmr/config"
	"github.com/bonedaddy/wxmr/db"
	"github.com/bonedaddy/wxmr/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	reserveContractAddress string
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
	cfg := config.DefaultConfig()
	t.Cleanup(func() {
		os.Remove(cfg.Database.Path)
		os.Remove(cfg.Logger.FilePath)
	})
	srv, err := NewService(ctx, config.DefaultConfig())

	// start the server
	go srv.Serve(ctx)

	// start the hacky miner
	go ethMine(t, srv)

	// prepare the test deposit
	srv.db.RegisterDeposit(addr1, id1)

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
	if resp, err := srv.mc.Transfer(rpc.TransferOpts{
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

	time.Sleep(time.Second * time.Duration(devConfirmationCount*5))

	mint, err := srv.db.GetMint(id1)
	require.NoError(t, err)
	require.Equal(t, db.Confirmed, mint.State)

	tx, err := srv.mc.GetTxID(srv.walletName, tx1)
	require.NoError(t, err)

	bal, err := srv.rsrv.BalanceOf(nil, common.HexToAddress(ethAddress))
	require.NoError(t, err)
	require.Equal(t, new(big.Int).SetUint64(tx.Transfer.Amount), bal)

	msg := time.Now().String()

	data, err = json.Marshal(&RequestReserveProof{Message: msg})
	require.NoError(t, err)

	req, err = http.NewRequest("GET", callURL+"/reserve_proof", bytes.NewReader(data))
	require.NoError(t, err)

	resp, err = hc.Do(req)
	require.NoError(t, err)

	data, err = ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var proof ResponseReserveProof
	require.NoError(t, json.Unmarshal(data, &proof))

	proofCheck, err := srv.mc.CheckReserveProof(srv.walletName, proof.ReserveAddress, msg, proof.Signature)
	require.NoError(t, err)
	t.Logf("reserve proof: %#v\n", proofCheck)
	cancel()
	// defer wait closed so its the last defer on the stack before exit
	srv.WaitClosed()
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
