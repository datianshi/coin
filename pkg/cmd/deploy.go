package cmd

import (
	"context"
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
	DeployCmd.Flags().StringVar(&clientURL, "client-url", "http://127.0.0.1:7545", "etheum client url, default to be ")
	DeployCmd.Flags().StringVar(&privateKey, "private-key", "", "private key in hex format")
	DeployCmd.MarkFlagRequired("private-key")
	DeployCmd.Flags().Int64Var(&value, "value", 0, "value to transfer along with the transaction")
	DeployCmd.Flags().Int64Var(&gasPrice, "gas-price", 1000000, "gas price for the transaction")
	DeployCmd.Flags().Uint64Var(&gasLimit, "gas-limit", 3000000, "gas limit for the transaction")
}

var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy --client-url [CLIENT_URL] --private-key [PRIVATE_KEY]",
	Run: func(cmd *cobra.Command, args []string) {
		addr, err := deploy()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Successfully deployed the contract to %s with address: %s\n", clientURL, addr.Hex())
	},
}

func deploy() (*common.Address, error) {
	client, err := ethclient.Dial(clientURL)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(value) // in wei
	auth.GasLimit = gasLimit       // in units
	auth.GasPrice = big.NewInt(gasPrice)

	address, _, _, err := contracts.DeployCoin(auth, client)
	if err != nil {
		return nil, err
	}

	return &address, err

}
