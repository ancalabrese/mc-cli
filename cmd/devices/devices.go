package devices

import (
	"fmt"

	"github.com/ancalabrese/mc-cli/mc/config"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/cobra"
)

func NewDevicesCommand(c *config.Config, l hclog.Logger) *cobra.Command {
	var devicesCmd = &cobra.Command{
		Use:   "devices",
		Short: "Manage your corporate devices",
		Long:  "Access your Mobicontrol devices information, run actions, check device policies and more.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("Method not implemented yet.")
		},
	}
	return devicesCmd
}
