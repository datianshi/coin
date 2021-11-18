package cmd

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/datianshi/coin/pkg/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var BalanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "balance",
	Run: func(cmd *cobra.Command, args []string) {
		var amount *big.Int
		var err error
		if amount, err = balance(); err != nil {
			panic(err)
		}
		fmt.Printf("The balance is %s\n", amount)
	},
}

func balance() (*big.Int, error) {
	var client *ethclient.Client
	var coin *contracts.Coin
	var key *ecdsa.PrivateKey
	var err error
	var amount *big.Int
	if client, err = ethclient.Dial(clientURL); err != nil {
		return nil, err
	}

	if key, err = crypto.HexToECDSA(privateKey); err != nil {
		return nil, err
	}

	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("invalid key")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	if coin, err = contracts.NewCoin(common.HexToAddress(contractAddress), client); err != nil {
		return nil, err
	}

	if amount, err = coin.Balances(&bind.CallOpts{}, address); err != nil {
		return nil, err
	}
	return amount, err
}
