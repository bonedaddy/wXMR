package service

// ResponseGetDepositAddress returns an address to use to deposit monero into
type ResponseGetDepositAddress struct {
	Address   string `json:"address"`
	PaymentID string `json:"payment_id"`
}

// RequestMint is used to request a minting of wXMR
type RequestMint struct {
	Address    string `json:"address"`
	PaymentID  string `json:"payment_id"`
	EthAddress string `json:"eth_address"` // the address to send the minted wXMR to
}

// RequestReserveProof is used to request a proof that all of the wallets balance is disposable
type RequestReserveProof struct {
	Message string `json:"message"` // the message to sign
}

type ResponseReserveProof struct {
	ReserveAddress string `json:"reserve_address"`
	Signature      string `json:"signature"`
}
