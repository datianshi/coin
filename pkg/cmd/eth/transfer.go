package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

func init() {
	TransferCmd.PersistentFlags().StringVar(&clientURL, "client-url", "http://127.0.0.1:7545", "etheum client url, default to be ")
	TransferCmd.PersistentFlags().StringVar(&privateKey, "private-key", "", "private key in hex format")
	TransferCmd.MarkFlagRequired("private-key")
	TransferCmd.PersistentFlags().Int64Var(&gasPrice, "gas-price", 1000000, "gas price for the transaction")
	TransferCmd.PersistentFlags().Uint64Var(&gasLimit, "gas-limit", 3000000, "gas limit for the transaction")
	TransferCmd.Flags().StringVar(&receiver, "receiver", "", "receiver address in hex format")
	TransferCmd.MarkFlagRequired("receiver")
	TransferCmd.Flags().Int64Var(&value, "value", 0, "eth value to transfer")
	TransferCmd.MarkFlagRequired("value")
}

var TransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "transfer --client-url [CLIENT_URL] --private-key [PRIVATE_KEY] --value [VALUE] --receiver [RECEIVER_ADDRESS]",
	Run: func(cmd *cobra.Command, args []string) {
		if err := transfer(); err != nil {
			panic(err)
		}
		fmt.Printf("Send to %s %d wei successfully", receiver, value)
	},
}

func transfer() error {

	client, err := ethclient.Dial(clientURL)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return err
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(receiver), big.NewInt(value), gasLimit, big.NewInt(gasPrice), nil)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

	if err != nil {
		return err
	}
	return client.SendTransaction(context.Background(), signedTx)
}
