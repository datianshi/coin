package analysis

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/spf13/cobra"
)

var address string
var fullTransaction bool

func init() {
	blockCmd.Flags().StringVar(&address, "address", "latest", "block address")
	blockCmd.Flags().BoolVar(&fullTransaction, "full-transaction", false, "full transaction blocks")

}

var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "block --address [ADDRESS] --full-transaction",
	Run: func(cmd *cobra.Command, args []string) {
		var data []byte
		var err error
		if data, err = showBlock(); err != nil {
			panic(err)
		}
		fmt.Printf(string(data))
	},
}

func showBlock() ([]byte, error) {
	var lastBlock interface{}
	client, err := rpc.Dial(clientURL)
	if err != nil {
		return nil, err
	}
	err = client.Call(&lastBlock, "eth_getBlockByNumber", address, fullTransaction)
	if err != nil {
		return nil, err
	}
	return json.Marshal(lastBlock)
}
