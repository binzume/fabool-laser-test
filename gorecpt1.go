package gorecpt1

// recpt1: https://github.com/stz2012/recpt1

import (
	"os/exec"
	"strings"
	"time"
	"fmt"
)

const DefaultCmd="recpt1"

type Recpt1 struct {
	Recpt1Cmd string
	Dev string
	RecOptions []string
}


func New(device string) *Recpt1 {
	return &Recpt1{Recpt1Cmd: DefaultCmd, Dev: device, RecOptions: []string{"--b25", "--strip"}}
}

func (r *Recpt1) GetChannels() (chs []string, err error) {
	cmd := exec.Command(r.Recpt1Cmd)
	result, err := cmd.Output()
	if err != nil {
		return []string{}, err
	}
	return strings.Split(strings.TrimSpace(string(result)),"\n"), nil
}

func (r *Recpt1) RecordCmd(ch string, ducation time.Duration, file string) string {
	return fmt.Sprintf("%s --dev %s %s %s %d %s", r.Recpt1Cmd, r.Dev, strings.Join(r.RecOptions, " "), ch, int(ducation.Seconds()), file)
}

