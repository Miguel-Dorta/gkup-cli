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
		Run: getSnapList,
	}
)

func init() {
	cmdRoot.AddCommand(cmdList)
}

func getSnapList(_ *cobra.Command, _ []string) {
	var snaps map[string][]int64
	out := gkup.Exec("-action=LIST", "-repo="+repoPath)
	if err := json.Unmarshal(out, &snaps); err != nil {
		log.Criticalf("error parsing snapshots: %s", err)
	}

	snapList := make([]string, len(snaps))
	for k := range snaps {
		snapList = append(snapList, k)
	}
	sort.Strings(snapList)

	for _, s := range snapList {
		timeList := snaps[s]
		sort.Slice(timeList, func(i, j int) bool { return timeList[i] < timeList[j] })

		if s == "" {
			s = "[no-name]"
		}
		fmt.Println(s)

		for _, i := range timeList {
			t := time.Unix(i, 0).Local()
			fmt.Printf("  - %04d/%02d/%02d %02d:%02d:%02d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
		}

		fmt.Println()
	}
}
