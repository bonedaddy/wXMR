package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	client "github.com/bonedaddy/wxmr/rpc"
	"github.com/monero-ecosystem/go-monero-rpc-client/wallet"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "wxmr"
	app.Usage = "wrapped monero cli"
	app.Commands = cli.Commands{
		&cli.Command{
			Name:  "new-address",
			Usage: "generate a new address under the account index",
			Action: func(c *cli.Context) error {
				cl, err := client.NewClient(c.String("wallet.rpc_address"))
				if err != nil {
					return err
				}
				addr, err := cl.NewAddress(c.String("wallet.name"), c.Uint64("account.index"))
				if err != nil {
					return err
				}
				fmt.Println("new address: ", addr)
				return cl.Close()
			},
		},
		&cli.Command{
			Name:  "new-wallet",
			Usage: "create a new monero wallet",
			Action: func(c *cli.Context) error {
				cl, err := client.NewClient(c.String("wallet.rpc_address"))
				if err != nil {
					return err
				}
				if err := cl.CreateWallet(c.String("wallet.name")); err != nil {
					return err
				}
				return cl.Close()
			},
		},
		&cli.Command{
			Name:    "convert-to-xmr",
			Aliases: []string{"ctxm"},
			Usage:   "converts an amount such as 0.1 to its corresponding XMR uint64 value",
			Action: func(c *cli.Context) error {
				fmt.Println(wallet.Float64ToXMR(c.Float64("value")))
				return nil
			},
		},
		&cli.Command{
			Name:    "get-churnable-addresses",
			Usage:   "returns all available churnable addresses",
			Aliases: []string{"gca"},
			Action: func(c *cli.Context) error {
				cl, err := client.NewClient(c.String("wallet.rpc_address"))
				if err != nil {
					return err
				}
				resp, err := cl.GetChurnableAddresses(c.String("wallet.name"), c.Uint64("churn.index"), c.Uint64("minimum.churn"))
				if err != nil {
					return err
				}
				fmt.Printf("%#v\n", resp)
				return cl.Close()
			},
		},
		&cli.Command{
			Name:  "address-balance",
			Usage: "retrieve balance for an address",
			Action: func(c *cli.Context) error {
				cl, err := client.NewClient(c.String("wallet.rpc_address"))
				if err != nil {
					return err
				}
				bal, err := cl.AddressBalance(c.String("wallet.name"), c.String("address"), c.Uint64("account.index"))
				if err != nil {
					return err
				}
				fmt.Println("balance: ", bal)
				return cl.Close()
			},
		},
		&cli.Command{
			Name:  "get-addresses",
			Usage: "returns all subaddresses underneath a given account index",
			Action: func(c *cli.Context) error {
				cl, err := client.NewClient(c.String("wallet.rpc_address"))
				if err != nil {
					return err
				}
				resp, err := cl.GetAddress(c.String("wallet.name"), c.Uint64("account.index"))
				if err != nil {
					return err
				}
				for _, r := range resp.Addresses {
					fmt.Println("address index: ", r.AddressIndex)
					fmt.Println("address: ", r.Address)
					fmt.Println("used: ", r.Used)
				}
				return cl.Close()
			},
		},
		&cli.Command{
			Name:  "get-all-accounts",
			Usage: "return all known accounts, their indexes, and subaddresses",
			Action: func(c *cli.Context) error {
				cl, err := client.NewClient(c.String("wallet.rpc_address"))
				if err != nil {
					return err
				}
				resp, err := cl.GetAccounts(c.String("wallet.name"))
				if err != nil {
					return err
				}
				for _, r := range resp.SubaddressAccounts {
					fmt.Println("account index: ", r.AccountIndex)
					fmt.Println("account base address: ", r.BaseAddress)
					fmt.Println("balance: ", r.Balance)
					fmt.Println("unlocked blance: ", r.UnlockedBalance)
				}
				return cl.Close()
			},
		},
		&cli.Command{
			Name:  "transfer",
			Usage: "used to transfer funds to the given address",
			Action: func(c *cli.Context) error {
				cl, err := client.NewClient(c.String("wallet.rpc_address"))
				if err != nil {
					return err
				}
				parts := c.StringSlice("subaddr.indices")
				var indices []uint64
				if len(parts) > 0 {
					for _, part := range parts {
						indice, err := strconv.ParseUint(part, 0, 64)
						if err != nil {
							continue // todo: handle
						}
						indices = append(indices, indice)
					}
				}
				resp, err := cl.Transfer(client.TransferOpts{
					WalletName:     c.String("wallet.name"),
					Destinations:   map[string]uint64{c.String("dest.address"): wallet.Float64ToXMR(c.Float64("value"))},
					AccountIndex:   c.Uint64("account.index"),
					SubaddrIndices: indices,
					Priority:       client.RandomPriority(),
				})
				if err != nil {
					return err
				}
				fmt.Println("tx hash: ", resp.TxHash)
				return cl.Close()
			},
		},
		&cli.Command{
			Name:  "rescan",
			Usage: "rescan entire blockchain, potentially destructive",
			Action: func(c *cli.Context) error {
				cl, err := client.NewClient(c.String("wallet.rpc_address"))
				if err != nil {
					return err
				}
				if err := cl.Rescan(c.String("wallet.name")); err != nil {
					return err
				}
				return cl.Close()
			},
		},
		&cli.Command{
			Name:  "refresh",
			Usage: "refresh accounts, scanning for incoming transactions",
			Action: func(c *cli.Context) error {
				cl, err := client.NewClient(c.String("wallet.rpc_address"))
				if err != nil {
					return err
				}
				if err := cl.Refresh(c.String("wallet.name")); err != nil {
					return err
				}
				return cl.Close()
			},
		},
		&cli.Command{
			Name:  "sweep-all",
			Usage: "sweep all accounts, use with caution",
			Action: func(c *cli.Context) error {
				cl, err := client.NewClient(c.String("wallet.rpc_address"))
				if err != nil {
					return err
				}
				resp, err := cl.SweepAll(client.TransferOpts{
					WalletName:   c.String("wallet.name"),
					AccountIndex: c.Uint64("account.index"),
					Destinations: map[string]uint64{c.String("dest.address"): wallet.Float64ToXMR(c.Float64("value"))},
					Priority:     client.RandomPriority(),
				})
				if err != nil {
					return err
				}
				fmt.Printf("%#v\n", resp)
				return cl.Close()
			},
		},
		&cli.Command{
			Name:  "sweep-dust",
			Usage: "sweeps all dust, use with caution",
			Action: func(c *cli.Context) error {
				cl, err := client.NewClient(c.String("wallet.rpc_address"))
				if err != nil {
					return err
				}
				resp, err := cl.SweepDust(c.String("wallet.name"))
				if err != nil {
					return err
				}
				fmt.Printf("%#v\n", resp)
				return cl.Close()
			},
		},
		&cli.Command{
			Name:  "mining",
			Usage: "mining related commands",
			Subcommands: cli.Commands{
				&cli.Command{
					Name:  "start",
					Usage: "start mining depositing funds into the given wallet",
					Action: func(c *cli.Context) error {
						cl, err := client.NewClient(c.String("wallet.rpc_address"))
						if err != nil {
							return err
						}
						if err := cl.StartMining(c.String("wallet.name"), c.Uint64("threads")); err != nil {
							return err
						}
						return cl.Close()
					},
					Flags: []cli.Flag{
						&cli.Uint64Flag{
							Name:  "threads",
							Usage: "number of threads to use for mining",
							Value: 2,
						},
					},
				},
				&cli.Command{
					Name:  "stop",
					Usage: "stop mining",
					Action: func(c *cli.Context) error {
						cl, err := client.NewClient(c.String("wallet.rpc_address"))
						if err != nil {
							return err
						}
						if err := cl.StopMining(c.String("wallet.name")); err != nil {
							return err
						}
						return cl.Close()
					},
				},
			},
		},
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "db.path",
			Usage: "path to sqlite database",
			Value: "mychurnerno_db",
		},
		&cli.StringFlag{
			Name:    "wallet.name",
			Aliases: []string{"wn"},
			Usage:   "the wallet to use for churning",
			Value:   "testnetwallet123",
		},
		&cli.StringFlag{
			Name:    "wallet.rpc_address",
			Aliases: []string{"wrpc"},
			Usage:   "the endpoint address of the monero-wallet-rpc server",
			Value:   "http://127.0.0.1:6061/json_rpc",
		},
		&cli.StringFlag{
			Name:    "dest.address",
			Aliases: []string{"da"},
			Usage:   "destination address to send funds to",
			Value:   "BhJQR4hu54wAqx9iRZZv5Y1UcTV6qgH52ULy5UNpEn7B7HVT2jpmAttf1k7mARTVWASvZkvajTk2NT5c2x3JHmojB5BDrFV",
		},
		&cli.StringFlag{
			Name:  "address",
			Usage: "the address to lookup",
		},
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"cfg"},
			Value:   "mychurnero.yml",
		},
		&cli.Uint64Flag{
			Name:    "account.index",
			Aliases: []string{"ai"},
			Usage:   "account index to use",
			Value:   0,
		},
		&cli.Uint64Flag{
			Name:    "minimum.churn",
			Aliases: []string{"mc"},
			Usage:   "minimum amount to churn from",
			Value:   1,
		},
		&cli.Uint64Flag{
			Name:  "churn.index",
			Usage: "account index to use for churning to",
			Value: 1,
		},
		&cli.StringSliceFlag{
			Name:  "subaddr.indices",
			Usage: "specify one or more subaddress indices to use",
		},
		&cli.Float64Flag{
			Name:  "value",
			Usage: "value to use generally for transfers",
			Value: 0.1,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
