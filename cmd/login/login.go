package login

import (
	"github.com/spf13/cobra"
)

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "log the CLI into Mobicontrol",
	Run: func(cmd *cobra.Command, args []string) {
		println(secret)
	},
}

var hostname, clientId, secret string

func init() {
	LoginCmd.Flags().StringVarP(
		&clientId,
		"clientId",
		"c",
		"",
		"add your Mobicontrol API client ID (required)")

	LoginCmd.Flags().StringVarP(
		&secret,
		"secret",
		"s",
		"",
		"add your Mobicontrol API client secret (required)")

	LoginCmd.Flags().StringVarP(
		&hostname,
		"host",
		"n",
		"",
		"add your Mobicontrol server host name, ie: s0000.mobicontrolcloud.com (required)")

	LoginCmd.MarkFlagRequired("clientId")
	LoginCmd.MarkFlagRequired("secret")
	LoginCmd.MarkFlagRequired("host")
}
