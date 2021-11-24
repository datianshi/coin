package main

import "github.com/datianshi/coin/pkg/cmd/analysis"

type Result struct {
	Result string
}

func main() {
	// var lastBlock string
	// // client, err := rpc.Dial("http://127.0.0.1:7545")
	// if err != nil {
	// 	panic(err)
	// }
	// err = client.Call(&lastBlock, "eth_getBalance", "0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8", "latest")
	// if err != nil {
	// 	panic(err)
	// }
	// value, err := hexutil.DecodeBig(string(lastBlock))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(value.String())

	analysis.AnalyzeCmd.Execute()

}
