package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	mclient "github.com/bonedaddy/wxmr/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
)

type Service struct {
	walletName string
	srv        *http.Server
	mc         *mclient.Client
	ec         *ethclient.Client
	auth       *bind.TransactOpts
}

func NewService(listenAddr, walletName string, ec *ethclient.Client, auth *bind.TransactOpts) (*Service, error) {
	srv := &Service{
		walletName: walletName,
		srv: &http.Server{
			Addr: listenAddr,
		},
	}
	r := chi.NewRouter()
	r.Get("/reserve_proof", srv.getReserveProof)
	r.Get("/deposit_address", srv.getDepositAddress)
	r.Post("/mint", srv.mint)
	return srv, nil
}

func (s *Service) getReserveProof(w http.ResponseWriter, r *http.Request) {}

func (s *Service) getDepositAddress(w http.ResponseWriter, r *http.Request) {
	pinfo, err := s.mc.NewIntegratedAddress(s.walletName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := &ResponseGetDepositAddress{
		Address:   pinfo.IntegratedAddress,
		PaymentID: pinfo.PaymentID,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.ServeContent(w, r, "", time.Time{}, bytes.NewReader(data))
}

func (s *Service) mint(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var req RequestMint
	if err := json.Unmarshal(data, &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("%+v\n", req)
	// TODO: queue processing the payment to trigger a mint
}
