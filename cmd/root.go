package cmd

import (
	"errors"
	"fmt"
	"os"

	"dilu/cmd/gen"
	"dilu/cmd/start"
	"dilu/cmd/version"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:          "dilu",
		Short:        "dilu",
		Long:         `dilu`,
		SilenceUsage: true,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				tip()
				return errors.New("requires at least one arg")
			}
			return nil
		},
		PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
		Run: func(cmd *cobra.Command, args []string) {
			tip()
		},
	}
)

func tip() {
	usageStr := `欢迎使用 dilu 查看命令：dilu --help`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(start.StartCmd)
	rootCmd.AddCommand(gen.GenCmd)
	rootCmd.AddCommand(version.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
