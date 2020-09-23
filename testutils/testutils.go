package testutils

import (
	"context"
	"math/big"
	"testing"

	"github.com/bonedaddy/wxmr/bindings/reserve"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	OneEthInWei = big.NewInt(1000000000000000000)
)

type Blockchain struct {
	auth *bind.TransactOpts
	*backends.SimulatedBackend
}

// NewBlockchain is used to generate a simulated blockchain
func NewBlockchain(t *testing.T) *Blockchain {
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(key)
	// https://medium.com/coinmonks/unit-testing-solidity-contracts-on-ethereum-with-go-3cc924091281
	balance := new(big.Int).Mul(big.NewInt(999999999999999), big.NewInt(999999999999999))
	gAlloc := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: balance},
	}
	sim := backends.NewSimulatedBackend(gAlloc, 8000000)
	return &Blockchain{auth, sim}
}

// DeployReserve is used to deploy the wXMR reserve contract on a test blockchain
func (b *Blockchain) DeployReserve() (common.Address, *reserve.Reserve, error) {
	addr, tx, contract, err := reserve.DeployReserve(b.auth, b, b.auth.From.Bytes())
	if err != nil {
		return common.Address{}, nil, err
	}
	b.Commit()
	if _, err := bind.WaitDeployed(context.Background(), b, tx); err != nil {
		return common.Address{}, nil, err
	}
	return addr, contract, err
}

// Mint is used to mint wXMR tokens
func (b *Blockchain) Mint(xmr *reserve.Reserve, recipient common.Address, amount *big.Int) error {
	tx, err := xmr.Mint(b.auth, recipient, amount)
	if err != nil {
		return err
	}
	b.Commit()
	_, err = bind.WaitMined(context.Background(), b, tx)
	return err
}

func (b *Blockchain) Burn(xmr *reserve.Reserve, amount *big.Int, destAddressHash [32]byte) error {
	tx, err := xmr.Burn(b.auth, amount, destAddressHash)
	if err != nil {
		return err
	}
	b.Commit()
	_, err = bind.WaitMined(context.Background(), b, tx)
	return err
}

func (b *Blockchain) Auth() *bind.TransactOpts {
	return b.auth
}
