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