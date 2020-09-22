package service

import (
	"log"
	"time"
)

func (s *Service) handleNewMint(paymentID, ethAddress string) {
	var confirmed bool
	for {
		resp, err := s.mc.GetPayments(s.walletName, paymentID)
		if err != nil {
			goto SLEEP
		}
		if len(resp.Payments) < 0 {
			goto SLEEP
		}
		if len(resp.Payments) > 1 {
			log.Println("ERROR: more than 2 payments found to this address")
		}
		confirmed, err = s.mc.TxConfirmed(s.walletName, resp.Payments[0].TxHash)
		if err != nil {
			log.Println("failed to check tx confirmation status: ", err)
		}
		if !confirmed {
			log.Println("tx is not yet confirmed")
		}
		if confirmed {
			// TODO: handle confirmation
			log.Println("tx is confirmed")
			return
		}
		goto SLEEP
	SLEEP:
		time.Sleep(time.Minute * 2)
	}
}
