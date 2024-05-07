package login

import (
	"fmt"

	"github.com/ancalabrese/mc-cli/mc/auth"
	"github.com/ancalabrese/mc-cli/mc/config"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/cobra"
)

func NewLoginCmd(c *config.Config, l hclog.Logger) *cobra.Command {
	var callback, hostname, clientId, secret string
	var loginCmd = &cobra.Command{
		Use:   "login",
		Short: "log the CLI into Mobicontrol",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := config.NewConfig(l)

			if cmd.Flags().NFlag() == 0 && (c.Host.ClientId == "" ||
				c.Host.ClientSecret == "" ||
				c.Host.HostName == "" ||
				c.Host.CallbackURL == "") {
				// no args, check config file
				return fmt.Errorf("Authentication info required")
			}

			if cmd.Flags().NFlag() > 0 {
				c.Host.ClientId, _ = cmd.Flags().GetString("clientId")
				c.Host.ClientSecret, _ = cmd.Flags().GetString("secret")
				c.Host.HostName, _ = cmd.Flags().GetString("host")
				c.Host.CallbackURL, _ = cmd.Flags().GetString("callback")
			}

			return auth.NewAuthSession(cmd.Context(), c, l)
		},
	}

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

	loginCmd.Flags().StringVarP(
		&callback,
		"callback",
		"u",
		"http://localhost:8080",
		"the Oauth2 callback url")

	loginCmd.MarkFlagsRequiredTogether("clientId", "secret", "host")
	return loginCmd
}
