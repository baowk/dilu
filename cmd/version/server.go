package version

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// Version is set via -ldflags at build time.
var Version = "dev"

var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: "dilu version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Println(resolveVersion())
	return nil
}

func resolveVersion() string {
	if Version != "" && Version != "dev" {
		return Version
	}

	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "" && info.Main.Version != "(devel)" {
			return info.Main.Version
		}
	}

	if Version == "" {
		return "dev"
	}
	return Version
}
