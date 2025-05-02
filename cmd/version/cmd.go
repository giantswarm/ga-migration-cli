package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {

	cmd := &cobra.Command{
		Use:          "version",
		Short:        "Print the version of the application",
		Run:          run,
		SilenceUsage: true,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println("Version: ", "0.0.1")
}
