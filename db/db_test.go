package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDatabaste(t *testing.T) {
	var testPath = "test.db"
	db, err := New(testPath)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, db.Close())
		os.RemoveAll(testPath)
	})

	t.Run("Deposit", func(t *testing.T) {
		type args struct {
			address   string
			paymentID string
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{"1-OK", args{"1", "1"}, false},
			{"1-Fail", args{"1", "1"}, true},
			{"2-OK", args{"2", "2"}, false},
			{"2-Fail", args{"2", "2"}, true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := db.RegisterDeposit(tt.args.address, tt.args.paymentID)
				if (err != nil) != tt.wantErr {
					t.Errorf("RegisterDeposit() err %v, wantErr %v", err, tt.wantErr)
				}
				dep, err := db.GetDeposit(tt.args.address)
				require.NoError(t, err)
				if !tt.wantErr {
					require.Equal(t, dep.Address, tt.args.address)
					require.Equal(t, dep.PaymentID, tt.args.paymentID)
					require.False(t, dep.Confirmed)
				}

				require.NoError(t, db.ConfirmDeposit(tt.args.address))

				dep, err = db.GetDeposit(tt.args.address)
				require.NoError(t, err)
				if !tt.wantErr {
					require.True(t, dep.Confirmed)
				}
			})
		}
	})

	t.Run("Mint", func(t *testing.T) {
		type args struct {
			paymentID  string
			ethAddress string
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{"1-OK", args{"1", "1"}, false},
			{"1-Fail", args{"1", "1"}, true},
			{"2-OK", args{"2", "2"}, false},
			{"2-Fail", args{"2", "2"}, true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := db.NewMint(tt.args.paymentID, tt.args.ethAddress)
				if (err != nil) != tt.wantErr {
					t.Errorf("RegisterDeposit() err %v, wantErr %v", err, tt.wantErr)
				}
				mint, err := db.GetMint(tt.args.paymentID)
				require.NoError(t, err)
				if !tt.wantErr {
					require.Equal(t, Pending, mint.State)

					require.NoError(t, db.SetMintState(tt.args.paymentID, Processing))

					mint, err = db.GetMint(tt.args.paymentID)
					require.NoError(t, err)
					require.Equal(t, Processing, mint.State)

					require.NoError(t, db.SetMintState(tt.args.paymentID, Confirmed))

					mint, err = db.GetMint(tt.args.paymentID)
					require.NoError(t, err)
					require.Equal(t, Confirmed, mint.State)

				}
			})
		}
	})
}
