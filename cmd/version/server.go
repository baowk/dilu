package version

import (
	"fmt"

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
	fmt.Println(Version)
	return nil
}
