package cmd

import (
	"fmt"

	"github.com/ancalabrese/mc-cli/cmd/devices"
	"github.com/ancalabrese/mc-cli/cmd/login"
	"github.com/ancalabrese/mc-cli/config"
	"github.com/ancalabrese/mc-cli/utils"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/cobra"
)

var (
	isVerbose bool = false
	l              = hclog.New(&hclog.LoggerOptions{
		Name:  "mc",
		Level: hclog.Error,
	})
	c = config.NewConfig(l)

	rootCmd = &cobra.Command{
		Use:   "mc",
		Short: "mc is a CLI for SOTI MobiControl",
		Long: "A very fast CLI tool that allows  IT Admins to quickly check and manage " +
			"corporate devices enrolled into SOTI MobiControl.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if isVerbose {
				l.SetLevel(hclog.Debug)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("no commands specified")
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&isVerbose, "verbose", "v", false, "enable verbose logging")

	rootCmd.AddCommand(login.NewLoginCmd(c, l), devices.NewDevicesCommand(c, l))
}

func Execute() {
	err := rootCmd.Execute()
	utils.Check(err)
}
