package service

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/bonedaddy/wxmr/bindings/reserve"
	"github.com/bonedaddy/wxmr/config"
	"github.com/bonedaddy/wxmr/db"
	"github.com/bonedaddy/wxmr/eth"
	mclient "github.com/bonedaddy/wxmr/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

var (
	devConfirmationCount  = 3
	prodConfirmationCount = 12
	dev                   = false
)

type Service struct {
	walletName string
	srv        *http.Server
	mc         *mclient.Client
	ec         *ethclient.Client
	auth       *bind.TransactOpts
	db         *db.Database
	rsrv       *reserve.Reserve
	logger     *zap.Logger
	ctx        context.Context
	closed     chan bool
	mux        sync.RWMutex
}

func NewService(ctx context.Context, cfg *config.Config) (*Service, error) {
	ec, err := cfg.EthRPC()
	if err != nil {
		return nil, err
	}
	auth, err := cfg.EthAuth()
	if err != nil {
		return nil, err
	}
	mc, err := cfg.XmrRPC()
	if err != nil {
		return nil, err
	}
	db, err := cfg.DB()
	if err != nil {
		return nil, err
	}
	logger, err := cfg.ZapLogger()
	if err != nil {
		return nil, err
	}
	var (
		contract *reserve.Reserve
		addr     common.Address
	)
	// allows deploying a contract before the service starts up
	if cfg.Ethereum.ReserveContractAddress == "" {
		addr, contract, err = eth.DeployReserveContract(ctx, auth, ec)
		if err != nil {
			return nil, err
		}
		logger.Info("reserve contract deployed", zap.String("contract.address", addr.String()))
	} else {
		contract, err = reserve.NewReserve(
			common.HexToAddress(cfg.Ethereum.ReserveContractAddress),
			ec,
		)
		if err != nil {
			return nil, err
		}
	}

	srv := &Service{
		walletName: cfg.Monero.WalletName,
		srv: &http.Server{
			Addr: cfg.Service.ListenAddr,
		},
		mc:     mc,
		ec:     ec,
		auth:   auth,
		db:     db,
		rsrv:   contract,
		ctx:    ctx,
		closed: make(chan bool, 1),
		logger: logger,
	}
	r := chi.NewRouter()
	r.Get("/reserve_proof", srv.getReserveProof)
	r.Get("/deposit_address", srv.getDepositAddress)
	r.Post("/mint", srv.startMint)
	srv.srv.Handler = r
	return srv, nil
}

func (s *Service) WaitClosed() {
	<-s.closed
}

func (s *Service) Serve(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		s.srv.Close()
		s.ec.Close()
		s.mc.Close()
		s.db.Close()
		s.closed <- true
	}()
	return s.srv.ListenAndServe()
}

func (s *Service) getReserveProof(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var req RequestReserveProof
	if err := json.Unmarshal(data, &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err := s.mc.GetReserveProof(s.walletName, req.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	walletAddr, err := s.mc.GetWalletAddress(s.walletName) // todo: replace with config variable to reduce rpc usage
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err = json.Marshal(&ResponseReserveProof{Signature: resp.Signature, ReserveAddress: walletAddr})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, r, "", time.Time{}, bytes.NewReader(data))

}

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
