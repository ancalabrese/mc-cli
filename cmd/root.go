package cmd

import (
	"fmt"

	"github.com/ancalabrese/mc-cli/cmd/devices"
	"github.com/ancalabrese/mc-cli/cmd/login"
	"github.com/ancalabrese/mc-cli/mc/config"
	"github.com/ancalabrese/mc-cli/utils"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "mc",
	Short: "mc is a CLI for SOTI MobiControl",
	Long: "A very fast CLI tool that allows  IT Admins to quickly check and manage " +
		"corporate devices enrolled into SOTI MobiControl.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("no commands specified")
	},
}

func init() {
	loggerOptions := &hclog.LoggerOptions{
		Name:  "mc",
		Level: hclog.Debug,
	}

	l := hclog.New(loggerOptions)
	c := config.NewConfig(l)

	root.AddCommand(login.NewLoginCmd(c, l), devices.NewDevicesCommand(c, l))
}

func Execute() {
	err := root.Execute()
	utils.Check(err)
}
