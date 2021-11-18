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

func init() {
	// FundCmd.Flags().StringVar(&receiver, "receiver", "", "receiver address in hex format")
	// FundCmd.MarkFlagRequired("receiver")
	FundCmd.Flags().Int64Var(&amount, "amount", 0, "amount to fund")
	FundCmd.MarkFlagRequired("amount")
}

var FundCmd = &cobra.Command{
	Use:   "fund",
	Short: "fund --amount [AMOUNT]",
	Run: func(cmd *cobra.Command, args []string) {
		if err := fund(); err != nil {
			panic(err)
		}
		fmt.Printf("Successfully fund %s with %d\n", receiver, amount)
	},
}

func fund() error {
	var client *ethclient.Client
	var coin *contracts.Coin
	var key *ecdsa.PrivateKey
	var err error
	if client, err = ethclient.Dial(clientURL); err != nil {
		return err
	}
	if key, err = crypto.HexToECDSA(privateKey); err != nil {
		return err
	}
	auth := bind.NewKeyedTransactor(key)
	auth.Value = big.NewInt(value) // in wei
	auth.GasLimit = gasLimit       // in units
	auth.GasPrice = big.NewInt(gasPrice)

	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	if coin, err = contracts.NewCoin(common.HexToAddress(contractAddress), client); err != nil {
		return err
	}

	if _, err = coin.Mint(auth, fromAddress, big.NewInt(amount)); err != nil {
		return err
	}
	return nil
}
