package service

import (
	"net/http"

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

	return srv, nil
}

func (s *Service) getReserveProof(w http.ResponseWriter, r *http.Request) {}
