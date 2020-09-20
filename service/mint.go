package service

import (
	"log"
	"time"
)

func (s *Service) handleNewMint(paymentID, ethAddress string) {
	for {
		resp, err := s.mc.GetPayments(s.walletName, paymentID)
		if err != nil {
			goto SLEEP
		}
		log.Printf("payments: %#v\n", resp.Payments)
		goto SLEEP
	SLEEP:
		time.Sleep(time.Minute * 2)
	}
}
