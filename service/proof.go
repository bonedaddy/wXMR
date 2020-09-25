package service

import (
	"math/big"
	"time"

	"github.com/bonedaddy/wxmr/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"go.uber.org/zap"
)

func (s *Service) submitProof() {
	addr, err := s.mc.GetWalletAddress(s.walletName)
	if err != nil {
		s.logger.Error("failed to get wallet address", zap.Error(err))
		return
	}
	msg := time.Now().UTC().String()
	proof, err := s.mc.GetReserveProof(s.walletName, msg)
	if err != nil {
		s.logger.Error("failed to get resreve proof", zap.Error(err))
		return
	}
	height, err := s.mc.BlockHeight(s.walletName)
	if err != nil {
		s.logger.Error("failed to get block height", zap.Error(err))
		return
	}
	proofHash := utils.SumKeccak256([]byte(proof.Signature))
	tx, err := s.rsrv.PostReserveProof(
		s.auth,
		[]byte(msg),
		[]byte(addr),
		proofHash,
		new(big.Int).SetUint64(height.Height),
	)
	if err != nil {
		s.logger.Error("failed to post reserve proof", zap.Error(err))
		return
	}
	s.logger.Info("reserve proof posted", zap.String("tx.hash", tx.Hash().String()))
	if _, err := bind.WaitMined(s.ctx, s.ec, tx); err != nil {
		s.logger.Error("failed to wait for transaction to be mined", zap.Error(err))
		return
	}
	s.logger.Info("reserve proof transaction mined", zap.String("tx.hash", tx.Hash().String()))
	// todo: update database
}
