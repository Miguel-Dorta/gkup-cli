package cmd

import (
	"fmt"
	"github.com/Miguel-Dorta/gkup-cli/internal"
	"github.com/Miguel-Dorta/gkup-cli/pkg/gkup"
	"github.com/spf13/cobra"
)

var (
	cmdVersion = &cobra.Command{
		Use:                        "version",
		Short:                      "print version",
		Run: func(_ *cobra.Command, _ []string) {
			out := gkup.Exec("-action=VERSION")
			fmt.Printf("gkup-cli: %s\ngkup: %s\n", internal.Version, string(out))
		},
	}
)

func init() {
	cmdRoot.AddCommand(cmdVersion)
}
