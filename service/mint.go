package service

import (
	"log"
	"math/big"
	"time"

	"github.com/bonedaddy/wxmr/db"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

func (s *Service) handleNewMint(paymentID, ethAddress string) {
	logger := s.logger.With(
		zap.String("payment.id", paymentID),
		zap.String("eth.address", ethAddress),
	)
	logger.Info("processing new mint request")
	for {
		resp, err := s.mc.GetPayments(s.walletName, paymentID)
		if err != nil {
			logger.Error("failed to retrieve payments", zap.Error(err))
			return
		}
		if len(resp.Payments) < 0 {
			logger.Debug("no payments found, sleeping for 2 minutes")
			time.Sleep(time.Minute * 2)
			continue
		}
		logger = logger.With(zap.String("monero.tx_hash", resp.Payments[0].TxHash))
		logger.Debug("found payment transaction, waiting for confirmations")
		if len(resp.Payments) > 1 {
			logger.Warn("more than two transactions found for this payment id, we only process the first one")
		}
	CONFIRM_START:
		confirmed, err := s.mc.TxConfirmed(s.walletName, resp.Payments[0].TxHash)
		if err != nil {
			logger.Error("failed to check tx confirmation status", zap.Error(err))
			return
		}
		if !confirmed {
			logger.Debug("transaction not confirmed, sleepign for 2 minutes")
			time.Sleep(time.Minute * 2)
			goto CONFIRM_START
		} else {
			logger.Debug("monero tx confirmed, starting wrapped token minting")
			tx, err := s.rsrv.Mint(
				s.auth,
				common.HexToAddress(ethAddress),
				new(big.Int).SetUint64(resp.Payments[0].Amount),
			)
			if err != nil {
				logger.Error("failed to create mint transaction", zap.Error(err))
				return
			}
			logger = logger.With(zap.String("ethereum.tx_hash", tx.Hash().String()))
			if err := s.db.SetMintState(paymentID, db.Processing); err != nil {
				logger.Error("failed to update mint state", zap.Error(err))
				return
			}
			if err := s.db.SetMintTx(paymentID, tx.Hash().String()); err != nil {
				logger.Error("failed to update mint transaction hash", zap.Error(err))
				return
			}
			logger.Debug("waiting for mint transaction to be mined")
			receipt, err := bind.WaitMined(s.ctx, s.ec, tx)
			if err != nil {
				logger.Error("failed to wait for transaction to be mined", zap.Error(err))
				return
			}
			if receipt.Status != 1 {
				logger.Error("transaction encountered a failure")
				return
			}
			logger.Debug("waiting for mint transaction to confirm")
			var (
				currentConfirmations *big.Int
				confirmationsNeeded  = getRequiredConfirmations()
				blockMinedAt         = receipt.BlockNumber
				lastBlockChecked     *big.Int
			)
			number := s.getCurrentBlockNumber()
			if number == nil {
				return
			}
			if number.Cmp(blockMinedAt) == 1 {
				currentConfirmations = new(big.Int).Sub(number, blockMinedAt)
			}
			lastBlockChecked = number
			for {
				sleep()
				number := s.getCurrentBlockNumber()
				if number == nil {
					return
				}
				if number.Cmp(lastBlockChecked) == 0 {
					continue
				}
				currentConfirmations = new(big.Int).Sub(number, lastBlockChecked)
				if num := currentConfirmations.Cmp(confirmationsNeeded); num == 0 || num == 1 {
					logger.Debug("transaction confirmed")
					break
				}
			}
			if err := s.db.SetMintState(paymentID, db.Confirmed); err != nil {
				logger.Debug("failed to set mint state", zap.Error(err))
				return
			}
			logger.Debug("successfully minted tokens")
			log.Println("tokens minted and confirmed")
			return
		}
	}
}

func (s *Service) getCurrentBlockNumber() *big.Int {
	header, err := s.ec.HeaderByNumber(s.ctx, nil)
	if err != nil {
		log.Println("failed to get current block number: ", err.Error())
		return nil
	}
	return header.Number
}

func getRequiredConfirmations() *big.Int {
	if dev {
		return big.NewInt(int64(devConfirmationCount))
	}
	return big.NewInt(int64(prodConfirmationCount))
}

func sleep() {
	if dev {
		time.Sleep(time.Second * time.Duration(devConfirmationCount))
		return
	}
	time.Sleep(time.Minute)
}
