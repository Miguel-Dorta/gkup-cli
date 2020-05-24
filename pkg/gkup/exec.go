package gkup

import (
	"bytes"
	"github.com/Miguel-Dorta/gkup-cli/pkg/log"
	"io"
	"os"
	"os/exec"
	"os/signal"
)

var (
	gkupCorePath string
	stop chan os.Signal
)

func init() {
	path, ok := os.LookupEnv("GKUP_PATH")
	if ok {
		gkupCorePath = path
	} else {
		gkupCorePath = "gkup-core"
	}

	stop = make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
}

func ExecPrintingStatus(args ...string) {
	gkup := exec.Command(gkupCorePath, args...)
	gkup.Stdout = newStdout()
	gkup.Stderr = newStderr()
	stdin, err := gkup.StdinPipe()
	if err != nil {
		log.Criticalf("error creating stdin pipe: %s", err)
		return
	}

	if err := gkup.Start(); err != nil {
		log.Critical(err.Error())
		return
	}

	done := make(chan error)
	go func() {
		done<-gkup.Wait()
	}()

	select {
	case <-stop:
		if _, err := io.WriteString(stdin, "STOP"); err != nil {
			log.Criticalf("error sending stop instruction to gkup-core: %s", err)
			return
		}
		log.Critical((<-done).Error())
	case err := <-done:
		if err != nil {
			log.Critical(err.Error())
		}
	}
}

func Exec(args ...string) []byte {
	buf := new(bytes.Buffer)

	gkup := exec.Command(gkupCorePath, args...)
	gkup.Stdout = buf
	gkup.Stderr = newStderr()
	stdin, err := gkup.StdinPipe()
	if err != nil {
		log.Criticalf("error creating stdin pipe: %s", err)
		return nil
	}

	if err := gkup.Start(); err != nil {
		log.Critical(err.Error())
		return nil
	}

	done := make(chan error)
	go func() {
		done<-gkup.Wait()
	}()

	select {
	case <-stop:
		if _, err := io.WriteString(stdin, "STOP"); err != nil {
			log.Criticalf("error sending stop instruction to gkup-core: %s", err)
			return nil
		}
		log.Critical((<-done).Error())
		return nil
	case err := <-done:
		if err != nil {
			log.Critical(err.Error())
			return nil
		}
	}
	return buf.Bytes()
}
