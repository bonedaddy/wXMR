package rpc

import (
	"github.com/monero-ecosystem/go-monero-rpc-client/wallet"
)

// NewAccount is used to create a new account with an optional label
func (c *Client) NewAccount(walletName, label string) (*wallet.ResponseCreateAccount, error) {
	if err := c.OpenWallet(walletName); err != nil {
		return nil, err
	}
	return c.mw.CreateAccount(&wallet.RequestCreateAccount{
		Label: label,
	})
}

// NewAddress creates a new address under the given account index
func (c *Client) NewAddress(walletName string, accountIndex uint64) (string, error) {
	if err := c.OpenWallet(walletName); err != nil {
		return "", err
	}
	resp, err := c.mw.CreateAddress(&wallet.RequestCreateAddress{AccountIndex: accountIndex})
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	return resp.Address, nil
}

// GetAccounts returns all accounts under the wallet
func (c *Client) GetAccounts(walletName string) (*wallet.ResponseGetAccounts, error) {
	if err := c.OpenWallet(walletName); err != nil {
		return nil, err
	}
	return c.mw.GetAccounts(&wallet.RequestGetAccounts{})
}

// GetWalletAddress is a convenience function to return wallet address information
func (c *Client) GetWalletAddress(walletName string) (string, error) {
	resp, err := c.GetAddress(walletName, 0)
	if err != nil {
		return "", err
	}
	return resp.Address, nil
}

// GetAddress returns address information for a given account index optionally filtered by subaddress index
func (c *Client) GetAddress(walletName string, accountIndex uint64, addressIndex ...uint64) (*wallet.ResponseGetAddress, error) {
	if err := c.OpenWallet(walletName); err != nil {
		return nil, err
	}
	return c.mw.GetAddress(&wallet.RequestGetAddress{
		AccountIndex: accountIndex,
		AddressIndex: addressIndex,
	})
}

// AddressBalance returns the unlocked funds for the given address
// TODO(bonedaddy): accept account and subaddress index
// look up balance for the given address (not the wallet)
func (c *Client) AddressBalance(walletName string, address string, accountIndex uint64, addressIndex ...uint64) (uint64, error) {
	if err := c.OpenWallet(walletName); err != nil {
		return 0, err
	}
	resp, err := c.mw.GetBalance(&wallet.RequestGetBalance{AccountIndex: accountIndex, AddressIndices: addressIndex})
	if err != nil {
		return 0, err
	}
	for _, addr := range resp.PerSubaddress {
		if addr.Address == address {
			return addr.UnlockedBalance, nil
		}
	}
	return 0, nil
}

func (c *Client) NewIntegratedAddress(walletName string) (*wallet.ResponseMakeIntegratedAddress, error) {
	if err := c.OpenWallet(walletName); err != nil {
		return nil, err
	}
	return c.mw.MakeIntegratedAddress(&wallet.RequestMakeIntegratedAddress{})
}

func (c *Client) SplitIntegratedAddress(walletName string) (*wallet.ResponseSplitIntegratedAddress, error) {
	if err := c.OpenWallet(walletName); err != nil {
		return nil, err
	}
	return c.mw.SplitIntegratedAddress(&wallet.RequestSplitIntegratedAddress{})
}

func (c *Client) GetPayments(walletName, paymentID string) (*wallet.ResponseGetPayments, error) {
	if err := c.OpenWallet(walletName); err != nil {
		return nil, err
	}
	return c.mw.GetPayments(&wallet.RequestGetPayments{PaymentID: paymentID})
}

func (c *Client) GetBulkPayments(walletName, paymentID string) (*wallet.ResponseGetBulkPayments, error) {
	if err := c.OpenWallet(walletName); err != nil {
		return nil, err
	}
	return c.mw.GetBulkPayments(&wallet.RequestGetBulkPayments{PaymentIDs: []string{paymentID}})
}
