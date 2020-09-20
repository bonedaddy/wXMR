package db

import "gorm.io/gorm"

// Deposit represents an integrated address
// that can be used to deposit XMR into the reserve minting wXMR
type Deposit struct {
	gorm.Model
	Address   string `gorm:"unique"`
	PaymentID string `gorm:"unique"`
	Confirmed bool
}

// RegisterDeposit is used to register a new deposit address
func (d *Database) RegisterDeposit(address, paymentID string) error {
	return d.db.Create(&Deposit{
		Address:   address,
		PaymentID: paymentID,
		Confirmed: false,
	}).Error
}

func (d *Database) GetDeposit(address string) (*Deposit, error) {
	var dep Deposit
	return &dep, d.db.Model(&Deposit{}).Where("address = ?", address).First(&dep).Error
}

func (d *Database) ConfirmDeposit(address string) error {
	dep, err := d.GetDeposit(address)
	if err != nil {
		return err
	}
	dep.Confirmed = true
	return d.db.Model(dep).Update("confirmed", true).Error
}
