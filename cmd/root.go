package cmd

import (
	"github.com/ancalabrese/mc-cli/cmd/login"
	"github.com/ancalabrese/mc-cli/utils"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "mc",
	Short: "mc is a CLI for SOTI MobiControl",
	Long: "A very fast CLI tool that allows an IT Admin to quickly check and manage " +
		"corporate devices enrolled into SOTI MobiControl.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	root.AddCommand(login.LoginCmd)
}

func Execute() {
	err := root.Execute()
	utils.Check(err)
}
