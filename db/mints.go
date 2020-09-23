package db

import "gorm.io/gorm"

// MintState represents a state about a given mint
type MintState string

func (m MintState) String() string {
	return string(m)
}

var (
	// Pending is the initial state of a mint
	// when we are waiting for the monero deposit to confirm
	Pending = MintState("PENDING")
	// Processing is the secondary state of a mint
	// when the monero deposit has confirmed, and when
	// we've broadcasted the wXMR mint transaction
	Processing = MintState("PROCESSING")
	// Confirmed is the third and last state of a mint
	// when the monero deposit has confirmed, and
	// the wXMR mint transaction has confirmed
	Confirmed = MintState("CONFIRMED")
)

// Mint represents a given minting event
type Mint struct {
	gorm.Model
	PaymentID       string `gorm:"unique"`
	EthAddress      string
	MintTransaction string
	State           MintState
}

// NewMint creates a new mint event
func (d *Database) NewMint(paymentID, ethAddress string) error {
	return d.db.Create(&Mint{
		PaymentID:  paymentID,
		EthAddress: ethAddress,
		State:      Pending,
	}).Error
}

func (d *Database) SetMintState(paymentID string, state MintState) error {
	mint, err := d.GetMint(paymentID)
	if err != nil {
		return err
	}
	mint.State = state
	return d.db.Model(mint).Update("state", state).Error
}

func (d *Database) GetMint(paymentID string) (*Mint, error) {
	var mint Mint
	return &mint, d.db.Model(&Mint{}).Where("payment_id = ?", paymentID).First(&mint).Error
}

func (d *Database) SetMintTx(paymentID string, txHash string) error {
	mint, err := d.GetMint(paymentID)
	if err != nil {
		return err
	}
	mint.MintTransaction = txHash
	return d.db.Model(mint).Update("mint_transaction", txHash).Error
}
