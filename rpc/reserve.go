package rpc

import (
	"errors"
	"time"

	"github.com/monero-ecosystem/go-monero-rpc-client/wallet"
)

// GetReserveProof returns a reserve proof for the given wallet
func (c *Client) GetReserveProof(walletName string) (*wallet.ResponseGetReserveProof, error) {
	if err := c.OpenWallet(walletName); err != nil {
		return nil, err
	}
	return c.mw.GetReserveProof(&wallet.RequestGetReserveProof{
		All:     true,
		Message: time.Now().UTC().String(),
	})
}

// CheckReserveProof is used to validate a reserve proof
func (c *Client) CheckReserveProof(walletName, reserveAddress, message, signature string) error {
	if err := c.OpenWallet(walletName); err != nil {
		return err
	}
	resp, err := c.mw.CheckReserveProof(&wallet.RequestCheckReserveProof{
		Address:   reserveAddress,
		Message:   message,
		Signature: signature,
	})
	if err != nil {
		return err
	}
	if !resp.Good {
		return errors.New("invalid proof")
	}
	return nil
}
