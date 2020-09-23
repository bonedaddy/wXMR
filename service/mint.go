package service

import (
	"log"
	"math/big"
	"time"

	"github.com/bonedaddy/wxmr/db"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (s *Service) handleNewMint(paymentID, ethAddress string) {
	for {
		resp, err := s.mc.GetPayments(s.walletName, paymentID)
		if err != nil {
			log.Println("failed to get payments: ", err.Error())
			return
		}
		if len(resp.Payments) < 0 {
			time.Sleep(time.Minute * 2)
			continue
		}
		if len(resp.Payments) > 1 {
			log.Println("ERROR: more than 2 payments found to this address")
		}
		confirmed, err := s.mc.TxConfirmed(s.walletName, resp.Payments[0].TxHash)
		if err != nil {
			log.Println("failed to check tx confirmation status: ", err.Error())
			return
		}
		if !confirmed {
			log.Println("tx is not yet confirmed")
			time.Sleep(time.Minute * 2)
			continue
		}
		if confirmed {
			log.Println("monero received, minting coins")
			tx, err := s.rsrv.Mint(
				s.auth,
				common.HexToAddress(ethAddress),
				new(big.Int).SetUint64(resp.Payments[0].Amount),
			)
			if err != nil {
				log.Println("failed to mint tokens: ", err.Error())
				return
			}
			if err := s.db.SetMintState(paymentID, db.Processing); err != nil {
				log.Println("failed to update mint state: ", err.Error())
				return
			}
			if err := s.db.SetMintTx(paymentID, tx.Hash().String()); err != nil {
				log.Println("failed to update mint transaction hash: ", err.Error())
				return
			}
			receipt, err := bind.WaitMined(s.auth.Context, s.ec, tx)
			if err != nil {
				log.Println("failed to wait for transaction to be mined: ", err.Error())
				return
			}
			if receipt.Status != 1 {
				log.Println("transaction status indicates failure: ", err.Error())
				return
			}
			var (
				currentConfirmations *big.Int
				confirmationsNeeded  = getRequiredConfirmations()
				blockMinedAt         = receipt.BlockNumber
				lastBlockChecked     *big.Int
			)
			header, err := s.ec.HeaderByNumber(s.auth.Context, nil)
			if err != nil {
				log.Println("failed to get current block number: ", err.Error())
				return
			}
			if header.Number.Cmp(blockMinedAt) == 1 {
				currentConfirmations = new(big.Int).Sub(header.Number, blockMinedAt)
			}
			lastBlockChecked = header.Number
			for {
				time.Sleep(time.Minute)
				header, err := s.ec.HeaderByNumber(s.auth.Context, nil)
				if err != nil {
					log.Println("failed to get current block number: ", err.Error())
					return
				}
				if header.Number.Cmp(lastBlockChecked) == 0 {
					continue
				}
				currentConfirmations = new(big.Int).Sub(header.Number, lastBlockChecked)
				if num := currentConfirmations.Cmp(confirmationsNeeded); num == 0 || num == 1 {
					log.Println("tx confirmed")
					break
				}
			}
			if err := s.db.SetMintState(paymentID, db.Confirmed); err != nil {
				log.Println("failed to set mint state: ", err.Error())
				return
			}
			log.Println("tokens minted and confirmed")
		}
		goto SLEEP
	SLEEP:
		time.Sleep(time.Minute * 2)
	}
}

func getRequiredConfirmations() *big.Int {
	if dev {
		return big.NewInt(int64(devConfirmationCount))
	}
	return big.NewInt(int64(prodConfirmationCount))
}
