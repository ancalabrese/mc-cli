package devices

import (
	"github.com/ancalabrese/mc-cli/actions"
	"github.com/ancalabrese/mc-cli/mc/client"
	"github.com/ancalabrese/mc-cli/mc/config"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/cobra"
)

func NewDevicesCommand(c *config.Config, l hclog.Logger) *cobra.Command {
	var take, skip int
	var path, deviceId string
	var devicesCmd = &cobra.Command{
		Use:   "devices",
		Short: "Manage your corporate devices",
		Long:  "Access your Mobicontrol devices information, run actions, check device policies and more.",
		RunE: func(cmd *cobra.Command, args []string) error {
			mcClient, err := client.NewMcClient(cmd.Context(), c, l)
			if err != nil {
				return err
			}

			if deviceId != "" {
				d, err := actions.GetDeviceById(cmd.Context(), mcClient, deviceId, l)
				if err != nil {
					return err
				}
				println(d.DeviceName)
				return nil
			}
			t := actions.Take(take)
			s := actions.Skip(skip)
			p := actions.Path(path)

			devices, err := actions.GetDevices(cmd.Context(), mcClient, l, t, s, p)
			if err != nil {
				return err
			}

			for _, d := range devices {
				println(d.DeviceName)
			}
			return nil
		},
	}

	devicesCmd.Flags().IntVarP(&take, "take", "t", 150, "the number of devices to be returned, after skipping over the 'skip' count")
	devicesCmd.Flags().IntVarP(&skip, "skip", "s", 0, "input the first X (count) devices that should not be returned")

	devicesCmd.Flags().StringVarP(&path, "path", "p", "", "the path of the parent device group. ie. '\\\\My Company\\BYOD'")

	devicesCmd.Flags().StringVarP(&deviceId, "deviceId", "i", "", "the ID of the device you want to check")

	return devicesCmd
}
