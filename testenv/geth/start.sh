#! /bin/bash

geth --datadir=datadir --rpc --rpcaddr 'localhost' --rpcport 8545 --rpcapi 'personal,db,eth,net,web3,txpool,miner' --dev 2> /dev/null &