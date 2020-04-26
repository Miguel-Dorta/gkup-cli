package gkup

import (
	"bytes"
	"github.com/Miguel-Dorta/gkup-cli/pkg/log"
	"os"
	"os/exec"
)

var gkupCorePath string

func init() {
	path, ok := os.LookupEnv("GKUP_PATH")
	if ok {
		gkupCorePath = path
	} else {
		gkupCorePath = "gkup-core"
	}
}

func ExecPrintingStatus(args ...string) {
	gkup := exec.Command(gkupCorePath, args...)
	gkup.Stdout = newStdout()
	gkup.Stderr = newStderr()

	if err := gkup.Run(); err != nil {
		log.Critical(err.Error())
	}
}

func Exec(args ...string) []byte {
	buf := new(bytes.Buffer)

	gkup := exec.Command(gkupCorePath, args...)
	gkup.Stdout = buf
	gkup.Stderr = newStderr()

	if err := gkup.Run(); err != nil {
		log.Critical(err.Error())
	}
	return buf.Bytes()
}
