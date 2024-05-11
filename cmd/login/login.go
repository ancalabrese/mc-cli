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

			if cmd.Flags().NFlag() == 0 && (c.Api.ClientId == "" || c.Api.ClientSecret == "" ||
				c.Api.HostName == "" || c.Api.CallbackURL == "") {
				return fmt.Errorf("API client info required.")
			}

			if cmd.Flags().NFlag() > 0 {
				c.Api.ClientId, _ = cmd.Flags().GetString("clientId")
				c.Api.ClientSecret, _ = cmd.Flags().GetString("secret")
				c.Api.HostName, _ = cmd.Flags().GetString("host")
				c.Api.CallbackURL, _ = cmd.Flags().GetString("callback")
				c.Write()
			}

			_, err := auth.NewAuthSession(cmd.Context(), c, l)
			if err != nil {
				return err
			}

			return c.Write()
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
