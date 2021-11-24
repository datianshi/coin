package analysis

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/spf13/cobra"
)

var txAddress string

func init() {
	txCmd.Flags().StringVar(&txAddress, "address", "", "tx address")
	txCmd.MarkFlagRequired("address")

}

var txCmd = &cobra.Command{
	Use:   "tx",
	Short: "tx --address [ADDRESS]",
	Run: func(cmd *cobra.Command, args []string) {
		var data []byte
		var err error
		if data, err = transaction(); err != nil {
			panic(err)
		}
		fmt.Printf(string(data))
	},
}

type Transaction struct {
	To    string
	Input string
}

func transaction() ([]byte, error) {
	var transaction interface{}
	client, err := rpc.Dial(clientURL)
	if err != nil {
		return nil, err
	}
	err = client.Call(&transaction, "eth_getTransactionByHash", txAddress)
	if err != nil {
		return nil, err
	}
	return json.Marshal(transaction)
}
