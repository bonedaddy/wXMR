package db

import "gorm.io/gorm"

// Deposit represents an integrated address
// that can be used to deposit XMR into the reserve minting wXMR
type Deposit struct {
	gorm.Model
	Address   string
	PaymentID string `gorm:"unique"`
	Confirmed bool
}

// RegisterDeposit is used to register a new deposit address
func (d *Database) RegisterDeposit(address, paymentID string) error {
	return d.db.Create(&Deposit{
		Address:   address,
		PaymentID: paymentID,
	}).Error
}
