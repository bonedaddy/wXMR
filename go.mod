module github.com/bonedaddy/wxmr

go 1.14

replace github.com/monero-ecosystem/go-monero-rpc-client v0.0.0-20200124164006-0afb4abdfc3c => github.com/bonedaddy/go-monero-rpc-client v0.0.1-rc2

require (
	github.com/ethereum/go-ethereum v1.9.21
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/monero-ecosystem/go-monero-rpc-client v0.0.0-20200124164006-0afb4abdfc3c
	github.com/segmentio/ksuid v1.0.3
	github.com/stretchr/testify v1.6.1
	github.com/urfave/cli/v2 v2.2.0
)
