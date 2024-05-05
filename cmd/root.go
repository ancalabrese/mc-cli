package cmd

import (
	"fmt"

	"github.com/ancalabrese/mc-cli/cmd/login"
	"github.com/ancalabrese/mc-cli/utils"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "mc",
	Short: "mc is a CLI for SOTI MobiControl",
	Long: "A very fast CLI tool that allows an IT Admin to quickly check and manage " +
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

	root.AddCommand(login.NewLoginCmd(l))
}

func Execute() {
	err := root.Execute()
	utils.Check(err)
}
