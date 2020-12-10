package cmd

import (
	"fmt"
	"github.com/panda8z/istone/src/pkg/global"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Get version info",
	Example: "istone version",
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runVersionCmd()
	},
}

func runVersionCmd() error {
	fmt.Println(global.Version)
	return nil
}
