package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Miguel-Dorta/gkup-cli/pkg/gkup"
	"github.com/Miguel-Dorta/gkup-cli/pkg/log"
	"github.com/spf13/cobra"
	"sort"
	"time"
)

var (
	cmdList = &cobra.Command{
		Use:   "list",
		Short: "list snapshots",
	}
)

func init() {
	cmdRoot.AddCommand(cmdList)
	cmdList.Run = getSnapList
}

func getSnapList(_ *cobra.Command, _ []string) {
	if repoPath == "" {
		cmdList.Help()
		log.Critical("repository path not defined")
	}

	var snaps map[string][]int64
	out := gkup.Exec("-action=LIST", "-repo="+repoPath)
	if err := json.Unmarshal(out, &snaps); err != nil {
		log.Criticalf("error parsing snapshots: %s", err)
	}

	snapNameList := make([]string, 0, len(snaps))
	for k := range snaps {
		snapNameList = append(snapNameList, k)
	}
	sort.Strings(snapNameList)

	for _, snapName := range snapNameList {
		if snapName != "" {
			fmt.Println(snapName)
		} else {
			fmt.Println("[no-name]")
		}

		sort.Slice(snaps[snapName], func(i, j int) bool {return snaps[snapName][i] < snaps[snapName][j]})
		for _, snapTime := range snaps[snapName] {
			t := time.Unix(snapTime, 0).Local()
			fmt.Printf("  - %04d/%02d/%02d %02d:%02d:%02d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
		}
		fmt.Println()
	}
}
