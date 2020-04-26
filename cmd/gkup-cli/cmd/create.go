package cmd

import (
	"github.com/Miguel-Dorta/gkup-cli/pkg/gkup"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	hashAlgorithm string
	cmdCreate = &cobra.Command{
		Use:                        "create",
		Short:                      "creates a gkup repository in the path provided",
		Run: func(_ *cobra.Command, args []string) {
			gkup.ExecPrintingStatus("-action=CREATE", "-repo="+repoPath, "-hash-algorithm="+hashAlgorithm, "-v="+strconv.FormatBool(verbose))
		},
	}
)

func init() {
	cmdRoot.AddCommand(cmdCreate)
	cmdCreate.Flags().StringVarP(&hashAlgorithm, "hash-algorithm", "h", "sha256", "Set repository's hash algorithm")
}
