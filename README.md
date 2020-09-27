# wXMR

wXMR is an experimental bridge between Ethereum and Monero allowing you to wrap Monero into an ERC20 token on the Ethereum blockchain. It is essentially a Monero equivalent of wBTC, and other wrapped cryptocurrencies. It allows anyone holding monero to deposit it into a public reserve which is an M/N multisignature wallet. Subsequently anyone can burn wXMR, retrieving the underlying XMR, withdrawing it from the public reserve. Reserve proofs are periodically stored on a smart contract to provide a historical record of the reserve balance, and anyone can poll an API to retrieve a current and up-to-date reserve proof.

# roadmap

* MVP
  * Basic bridge allowing wrapping of XMR and unwrapping (burning) of wXMR [done]
  * Basic API to retrieve up to date reserve proof [done]
  * Basic API to retrieve all known reserve proofs [wip]
  * Basic API to retrieve reserve proof signature by it's hash [wip]
* Alpha 1
  * Move reserve wallet to multisignature wallet
* Alpha 2
  * Generate monero multisignature wallet using ethereum keys
* Alpha 3
  * Governance token to elect community managed reserves via stake in governance token using ETH or DAI bonds.

# testenv

To bootstrap a new testenv you can use the following commands:

```shell
$> make start-testenv # best to do this in a screen session
$> make # builds the cli
$> ./wxmr new-wallet # generates a new wallet
$> ./wmxr mining start # starts the mining proces
```


# monero-wallet-rpc

To start the wallet rpc service run:

```shell
$> monero-wallet-rpc --testnet --disable-rpc-login --prompt-for-password --rpc-bind-port 6061 --daemon-port 28081 --wallet-dir=.
```