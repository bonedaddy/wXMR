package client

import (
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
