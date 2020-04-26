package cmd

import (
	"github.com/Miguel-Dorta/gkup-cli/pkg/log"
	"github.com/spf13/cobra"
)

var (
	repoPath string
	verbose bool
	cmdRoot = &cobra.Command{
		Use:                        "gkup",
		Short:                      "gkup is a backup tool optimized to eliminate duplicate files",
	}
)

func init() {
	cmdRoot.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Set verbose output (may affect performance)")
	cmdRoot.PersistentFlags().StringVarP(&repoPath, "repo", "r", "", "Repository path")
}

func Execute() {
	if err := cmdRoot.Execute(); err != nil {
		log.Critical(err.Error())
	}
}
