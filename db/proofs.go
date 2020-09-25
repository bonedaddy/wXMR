package db

import "gorm.io/gorm"

// Proof are the official proofs submitted to the reserve contract
type Proof struct {
	gorm.Model
	MoneroHeight   uint64
	EthereumHeight int64
	ProofHash      string `gorm:"unique"` // the hash of the proof signature
	ProofSignature string
}

// NewProof stores a new proof in the database
func (d *Database) NewProof(
	moneroHeight uint64,
	ethereumHeight int64,
	proofHash, proofSignature string,
) error {
	return d.db.Create(&Proof{
		MoneroHeight:   moneroHeight,
		EthereumHeight: ethereumHeight,
		ProofHash:      proofHash,
		ProofSignature: proofSignature,
	}).Error
}

// GetProof searches for a proof by its hash
func (d *Database) GetProof(hash string) (*Proof, error) {
	var proof Proof
	return &proof, d.db.Model(
		&Proof{},
	).Where("proof_hash = ?", hash).First(&proof).Error
}

// GetProofs returns all known proofs
func (d *Database) GetProofs() ([]Proof, error) {
	var proofs []Proof
	return proofs, d.db.Model(&Proof{}).Find(&proofs).Error
}
