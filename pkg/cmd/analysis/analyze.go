package analysis

import "github.com/spf13/cobra"

var clientURL string

func init() {
	AnalyzeCmd.PersistentFlags().StringVar(&clientURL, "client-url", "http://127.0.0.1:7545", "etheum client url, default to be ")
	AnalyzeCmd.AddCommand(blockCmd)
	AnalyzeCmd.AddCommand(txCmd)
}

var AnalyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "analyze --client-url [CLIENT_URL]",
}
