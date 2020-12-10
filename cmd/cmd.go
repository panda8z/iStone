/*

*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/panda8z/istone/src/pkg/global"
	"os"

	"github.com/panda8z/istone/src/pkg/tools"
	"github.com/spf13/cobra"
)

var (
	configYml  string
	mode       string
	port       string
	traceStart bool
)

var rootCmd = &cobra.Command{
	Use:          "istone",
	Short:        "istone",
	SilenceUsage: true,
	Long:         `istone`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(tools.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(configCmd)
}

func tip() {
	usageStr := `欢迎使用 ` + tools.Green(`istone `+global.Version) + ` 可以使用 ` + tools.Red(`-h`) + ` 查看命令`
	usageStr1 := `也可以参考 http://doc.zhangwj.com/istone-site/guide/ksks.html 里边的【启动】章节`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
