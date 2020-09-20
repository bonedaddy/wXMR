package service

// ResponseGetDepositAddress returns an address to use to deposit monero into
type ResponseGetDepositAddress struct {
	Address   string `json:"address"`
	PaymentID string `json:"payment_id"`
}

// RequestMint is used to request a minting of wXMR
type RequestMint struct {
	PaymentID  string `json:"payment_id"`
	EthAddress string `json:"eth_address"` // the address to send the minted wXMR to
}
