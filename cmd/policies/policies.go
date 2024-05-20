package policies

import (
	"errors"

	"github.com/ancalabrese/mc-cli/client"
	"github.com/ancalabrese/mc-cli/config"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/cobra"
)

func NewPoliciesCommand(c *config.Config, l hclog.Logger) *cobra.Command {

	var profileCmd = &cobra.Command{
		Use:   "policies",
		Short: "Manage device policies",
		Long:  "Access information about any Mobicontrol device policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := client.NewMcClient(cmd.Context(), c, l)
			if err != nil {
				return err
			}

			return errors.New("function not implemented yet.")
		},
	}
	return profileCmd
}
