package login

import (
	"github.com/ancalabrese/mc-cli/mc/config"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/cobra"
)

var hostname, clientId, secret string
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "log the CLI into Mobicontrol",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			// no args, check config file
		}
		// create config and start auth

	},
}

func NewLoginCmd(c *config.Config, l hclog.Logger) *cobra.Command {
	loginCmd.Flags().StringVarP(
		&clientId,
		"clientId",
		"c",
		"",
		"add your Mobicontrol API client ID")

	loginCmd.Flags().StringVarP(
		&secret,
		"secret",
		"s",
		"",
		"add your Mobicontrol API client secret")

	loginCmd.Flags().StringVarP(
		&hostname,
		"host",
		"n",
		"",
		"add your Mobicontrol server host name, ie: s0000.mobicontrolcloud.com ")

	loginCmd.MarkFlagsRequiredTogether("clientId", "secret", "host")
	return loginCmd
}
