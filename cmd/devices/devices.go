package devices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ancalabrese/mc-cli/client"
	"github.com/ancalabrese/mc-cli/config"
	"github.com/ancalabrese/mc-cli/data"
	"github.com/ancalabrese/mc-cli/screen"
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

			devices := []*data.BaseDevice{}

			if deviceId != "" {
				d, err := mcClient.GetDeviceById(cmd.Context(), deviceId)
				if err != nil {
					return err
				}
				devices = append(devices, d)
			} else {
				t := mcClient.Take(take)
				s := mcClient.Skip(skip)
				p := mcClient.Path(path)

				devices, err = mcClient.GetDevices(cmd.Context(), t, s, p)
				if err != nil {
					return err
				}
			}

			if len(devices) == 0 {
				println("Feels so empty here! Start enrolling devices!")
			}

			var selection = 1
			if len(devices) > 1 {
				for i, d := range devices {
					fmt.Printf("#[%d] > %s - ID: %s\n", i+1, d.DeviceName, d.DeviceId)
				}

				r := bufio.NewReader(os.Stdin)

				for ok := true; ok; {
					fmt.Println("Enter the device # to check details:")
					input, err := r.ReadString('\n')
					if err != nil {
						return err
					}
					input = strings.TrimSuffix(input, "\n")

					selection, err = strconv.Atoi(input)
					if err != nil || selection-1 < 0 || selection-1 > len(devices) {
						fmt.Println("Enter a valid number.")
						continue
					}
					break
				}
			}

			p := screen.NewPrinter(os.Stdout)
			err = p.PrettyPrint(devices[selection-1])
			if err != nil {
				return err
			}
			return nil
		},
	}

	devicesCmd.Flags().IntVarP(&take, "take", "t", 150, "the number of devices to be returned, after skipping over the 'skip' count")
	devicesCmd.Flags().IntVarP(&skip, "skip", "s", 0, "input the first X (count) devices that should not be returned")

	devicesCmd.Flags().StringVarP(&path, "path", "p", "", "the path of the parent device group. ie. '\\\\My Company\\BYOD'")

	devicesCmd.Flags().StringVarP(&deviceId, "id", "i", "", "the ID of the device you want to check")

	devicesCmd.AddCommand(NewDeviceDeleteCmd(c, l))
	return devicesCmd
}
