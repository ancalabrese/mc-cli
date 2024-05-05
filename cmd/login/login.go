package login

import (
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/cobra"
)

var hostname, clientId, secret string
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "log the CLI into Mobicontrol",
	Run: func(cmd *cobra.Command, args []string) {
		println(secret)
	},
}

func NewLoginCmd(l hclog.Logger) *cobra.Command {
	loginCmd.Flags().StringVarP(
		&clientId,
		"clientId",
		"c",
		"",
		"add your Mobicontrol API client ID (required)")

	loginCmd.Flags().StringVarP(
		&secret,
		"secret",
		"s",
		"",
		"add your Mobicontrol API client secret (required)")

	loginCmd.Flags().StringVarP(
		&hostname,
		"host",
		"n",
		"",
		"add your Mobicontrol server host name, ie: s0000.mobicontrolcloud.com (required)")

	loginCmd.MarkFlagRequired("clientId")
	loginCmd.MarkFlagRequired("secret")
	loginCmd.MarkFlagRequired("host")

	return loginCmd
}
