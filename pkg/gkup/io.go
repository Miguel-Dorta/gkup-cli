package gkup

import (
	"bytes"
	"encoding/json"
	"github.com/Miguel-Dorta/gkup-cli/pkg/log"
	"io"
)

type status struct {
	Global struct {
		Current uint64 `json:"current"`
		Total   uint64 `json:"total"`
		Name    string `json:"name"`
	} `json:"global"`
	Partial struct {
		Current uint64 `json:"current"`
		Total   uint64 `json:"total"`
		Details string `json:"details"`
	} `json:"partial"`
}

type stdout struct {
	buf          *bytes.Buffer
	j            *json.Decoder
}

type stderr struct {
	buf *bytes.Buffer
}

const (
	globalFormat = "(%d/%d) %s"
	partialFormat = globalFormat + ": " + globalFormat
)

func newStdout() *stdout {
	b := new(bytes.Buffer)
	return &stdout{
		buf:          b,
		j:            json.NewDecoder(b),
	}
}

func newStderr() *stderr {
	return &stderr{buf: new(bytes.Buffer)}
}

func (w *stdout) Write(data []byte) (int, error) {
	w.buf.Write(data)
	for w.j.More() {
		var s status
		if err := w.j.Decode(&s); err != nil {
			return 0, err
		}
		w.print(&s)
	}
	return len(data), nil
}

func (w *stdout) print(s *status) {
	if s.Partial.Current == 0 {
		log.Infof(globalFormat, s.Global.Current, s.Global.Total, s.Global.Name)
	} else {
		log.Debugf(partialFormat, s.Global.Current, s.Global.Total, s.Global.Name, s.Partial.Current, s.Partial.Total, s.Partial.Details)
	}
}

func (w *stderr) Write(data []byte) (int, error) {
	w.buf.Write(data)
	for {
		s, err := w.buf.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			w.buf.WriteString(s)
			return len(data), nil
		}
		log.Error(s[:len(s)-1])
	}
}
