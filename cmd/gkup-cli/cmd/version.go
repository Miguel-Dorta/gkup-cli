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
			fmt.Printf("gkup-cli: %s\n", internal.Version)
			fmt.Printf("gkup-core: %s\n", gkup.Exec("-action=VERSION"))
		},
	}
)

func init() {
	cmdRoot.AddCommand(cmdVersion)
}
