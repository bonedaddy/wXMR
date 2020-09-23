package eth

import (
	"context"

	"github.com/bonedaddy/wxmr/bindings/reserve"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// DeployReserveContract is used to deploy the reserve contract
func DeployReserveContract(ctx context.Context, auth *bind.TransactOpts, ec *ethclient.Client) (common.Address, *reserve.Reserve, error) {
	addr, tx, contract, err := reserve.DeployReserve(auth, ec)
	if err != nil {
		return common.Address{}, nil, err
	}
	if _, err := bind.WaitDeployed(ctx, ec, tx); err != nil {
		return common.Address{}, nil, err
	}
	return addr, contract, nil
}
