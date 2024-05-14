package devices

import (
	"github.com/ancalabrese/mc-cli/client"
	"github.com/ancalabrese/mc-cli/config"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/cobra"
)

func NewDeviceDeleteCmd(c *config.Config, logger hclog.Logger) *cobra.Command {
	var deviceId string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a managed device",
		Long:  "Delete a managed device. Frees up a licence, and revokes control over the device.",
		RunE: func(cmd *cobra.Command, args []string) error {
			mcClient, err := client.NewMcClient(cmd.Context(), c, logger)
			if err != nil {
				return err
			}

			err = mcClient.DeleteDevice(cmd.Context(), deviceId)
			return err
		},
	}

	cmd.Flags().StringVarP(&deviceId, "deviceId", "i", "", "Specify the ID of the device to be deleted")
	cmd.MarkFlagRequired("deviceId")
	return cmd
}
