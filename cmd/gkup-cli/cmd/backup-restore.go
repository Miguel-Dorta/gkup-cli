package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Miguel-Dorta/gkup-cli/pkg/gkup"
	"github.com/Miguel-Dorta/gkup-cli/pkg/log"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"time"
)

var (
	snapName, snapTime string
	cmdBackup          = &cobra.Command{
		Use:   "backup",
		Short: "take a snapshot of the files provided",
		Run: func(_ *cobra.Command, args []string) {
			gkupArgs := []string{
				"-action=BACKUP",
				"-repo=" + repoPath,
				"-snapshot-name=" + snapName,
				"-v=" + strconv.FormatBool(verbose),
			}
			gkupArgs = append(gkupArgs, args...)
			gkup.ExecPrintingStatus(gkupArgs...)
		},
	}
	cmdRestore = &cobra.Command{
		Use:   "restore",
		Short: "restore a snapshot",
		Run: func(_ *cobra.Command, args []string) {
			if len(args) < 1 {
				log.Critical("destination path for restoring was not provided")
			}
			if len(args) > 1 {
				log.Critical("more than one destination path was provided")
			}
			gkup.ExecPrintingStatus(
				"-action=RESTORE",
				"-repo="+repoPath,
				"-snapshot-name="+snapName,
				"-snapshot-time="+strconv.FormatInt(findSnap(snapTime, snapName), 10),
				"-restore-destination="+args[0],
				"-v="+strconv.FormatBool(verbose))
		},
	}
)

func init() {
	cmdRoot.AddCommand(cmdBackup, cmdRestore)
	cmdBackup.Flags().StringVarP(&snapName, "snapshot-name", "n", "", "Set the snapshot name (multiple snapshots can be grouped under the same name)")
	cmdRestore.Flags().StringVarP(&snapName, "snapshot-name", "n", "", "Set the snapshot name (multiple snapshots can be grouped under the same name)")
	cmdRestore.Flags().StringVarP(&snapTime, "snapshot-time", "t", "", "Set the snapshot time. Must be in \"YYYY/MM/DD hh:mm:ss\" format and can be partial")
}

func findSnap(query, snapName string) int64 {
	var snapTimes map[int64]string
	{
		var snaps map[string][]int64
		out := gkup.Exec("-action=LIST", "-repo="+repoPath)
		if err := json.Unmarshal(out, &snaps); err != nil {
			log.Criticalf("error parsing snapshots: %s", err)
		}
		snapTimes = make(map[int64]string, len(snaps))

		iList := snaps[snapName]
		for _, i := range iList {
			t := time.Unix(i, 0).Local()
			snapTimes[i] = fmt.Sprintf("%04d/%02d/%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
		}
	}

	snapTimes = filterMapInt64String(snapTimes, func(_ int64, v string) bool { return strings.HasPrefix(v, query) })
	if len(snapTimes) > 1 {
		log.Criticalf("more than one snapshot matches the time provided: %+v", snapTimes)
	}
	for k, _ := range snapTimes {
		return k
	}
	log.Critical("no snapshot matches the time provided")
	return 0
}

func filterMapInt64String(m map[int64]string, filterFunc func(k int64, v string) bool) map[int64]string {
	result := make(map[int64]string, len(m))
	for k, v := range m {
		if filterFunc(k, v) {
			result[k] = v
		}
	}
	return result
}
