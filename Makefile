.PHONY: contracts
contracts:
		solc --bin --abi --optimize --optimize-runs 200 --hashes --devdoc --userdoc --pretty-json --output-dir=bin contracts/wXMR.sol

.PHONY: bindings
bindings:
	abigen --pkg wxmr --abi bin/wXMR.abi --bin bin/wXMR.bin --out bindings/wxmr/wxmr.go