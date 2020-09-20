package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bonedaddy/wxmr/db"
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
	db         *db.Database
}

func NewService(listenAddr, walletName string, mc *mclient.Client, ec *ethclient.Client, auth *bind.TransactOpts, database *db.Database) (*Service, error) {
	srv := &Service{
		walletName: walletName,
		srv: &http.Server{
			Addr: listenAddr,
		},
		mc:   mc,
		ec:   ec,
		auth: auth,
		db:   database,
	}
	r := chi.NewRouter()
	r.Get("/reserve_proof", srv.getReserveProof)
	r.Get("/deposit_address", srv.getDepositAddress)
	r.Post("/mint", srv.startMint)
	return srv, nil
}

func (s *Service) getReserveProof(w http.ResponseWriter, r *http.Request) {}

func (s *Service) getDepositAddress(w http.ResponseWriter, r *http.Request) {
	pinfo, err := s.mc.NewIntegratedAddress(s.walletName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := s.db.RegisterDeposit(pinfo.IntegratedAddress, pinfo.PaymentID); err != nil {
		http.Error(w, "failed to create deposit address", http.StatusInternalServerError)
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

func (s *Service) startMint(w http.ResponseWriter, r *http.Request) {
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
	if err := s.db.NewMint(req.PaymentID, req.EthAddress); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	go s.handleNewMint(req.PaymentID, req.EthAddress)
}
