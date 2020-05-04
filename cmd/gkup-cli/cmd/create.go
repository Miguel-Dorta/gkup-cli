package cmd

import (
	"github.com/Miguel-Dorta/gkup-cli/pkg/gkup"
	"github.com/Miguel-Dorta/gkup-cli/pkg/log"
	"github.com/Miguel-Dorta/gkup-core/pkg/hash"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	hashAlgorithm string
	cmdCreate     = &cobra.Command{
		Use:   "create",
		Short: "creates a gkup repository in the path provided",
	}
)

func init() {
	cmdRoot.AddCommand(cmdCreate)
	cmdCreate.Run = create
	cmdCreate.Flags().StringVarP(&hashAlgorithm, "hash-algorithm", "H", "sha256", "Set repository's hash algorithm")
}

func create(_ *cobra.Command, args []string) {
	if repoPath == "" {
		cmdCreate.Help()
		log.Critical("repository path not defined")
	}

	switch hashAlgorithm {
	case hash.MD5, hash.SHA1, hash.SHA256, hash.SHA512, hash.SHA3_256, hash.SHA3_512:
	default:
		log.Critical("invalid hash algorithm")
	}
	gkup.ExecPrintingStatus("-action=CREATE", "-repo="+repoPath, "-hash-algorithm="+hashAlgorithm, "-v="+strconv.FormatBool(verbose))
}
