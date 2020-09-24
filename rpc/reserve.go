package rpc

import (
	"errors"

	"github.com/monero-ecosystem/go-monero-rpc-client/wallet"
)

// GetReserveProof returns a reserve proof for the given wallet
func (c *Client) GetReserveProof(walletName, message string) (*wallet.ResponseGetReserveProof, error) {
	if err := c.OpenWallet(walletName); err != nil {
		return nil, err
	}
	return c.mw.GetReserveProof(&wallet.RequestGetReserveProof{
		All:     true,
		Message: message,
	})
}

// CheckReserveProof is used to validate a reserve proof
func (c *Client) CheckReserveProof(walletName, reserveAddress, message, signature string) (*wallet.ResponseCheckReserveProof, error) {
	if err := c.OpenWallet(walletName); err != nil {
		return nil, err
	}
	resp, err := c.mw.CheckReserveProof(&wallet.RequestCheckReserveProof{
		Address:   reserveAddress,
		Message:   message,
		Signature: signature,
	})
	if err != nil {
		return nil, err
	}
	if !resp.Good {
		return nil, errors.New("invalid proof")
	}
	return resp, nil
}
