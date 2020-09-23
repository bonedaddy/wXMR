.PHONY: build
build:
	go build -o wxmr ./cmd/wxmr
	
.PHONY: contracts
contracts:
		solc \
			--bin \
			--abi \
			--optimize \
			--optimize-runs 200 \
			--hashes \
			--devdoc \
			--userdoc \
			--pretty-json \
			--output-dir=bin \
			--overwrite \
			contracts/Reserve.sol

.PHONY: bindings
bindings:
	abigen --pkg reserve --abi bin/Reserve.abi --bin bin/Reserve.bin --out bindings/reserve/reserve.go

.PHONY: start-testenv
start-testenv:
	(cd testenv/monero ; bash start.sh)
