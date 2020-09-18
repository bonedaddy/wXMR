package eth

import (
	"context"
	"log"

	"github.com/bonedaddy/wxmr/bindings/reserve"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// Listener provides convenience functions around a reserve bindings
type Listener struct {
	*reserve.Reserve
}

// NewListener returns a new reserve contract listener
func NewListener(rsrv *reserve.Reserve) *Listener {
	return &Listener{rsrv}
}

// HandleCoinBurn processes coin burn events
func (l *Listener) HandleCoinBurn(ctx context.Context, output chan *reserve.ReserveCoinsBurned) error {
	ch := make(chan *reserve.ReserveCoinsBurned)
	sub, err := l.WatchCoinsBurned(&bind.WatchOpts{Context: ctx}, ch)
	if err != nil {
		return err
	}
	for {
		select {
		case err := <-sub.Err():
			log.Printf("subscription error: %s", err.Error())
			return err
		case ev := <-ch:
			log.Printf("new reserve burn detected: %+v\n", ev)
			output <- ev
		}
	}
}
