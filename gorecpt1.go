package gorecpt1

// recpt1: https://github.com/stz2012/recpt1

import (
	"os/exec"
	"strings"
	"time"
	"fmt"
)

const Recpt1="recpt1"

func GetChannels() (chs []string, err error) {
	cmd := exec.Command(Recpt1)
	result, err := cmd.Output()
	if err != nil {
		return []string{}, err
	}
	return strings.Split(strings.TrimSpace(string(result)),"\n"), nil
}

func RecordCmd(ch string, ducation time.Duration, file string, option[]string) string {
	return fmt.Sprintf("%s %s %s %d %s", Recpt1, strings.Join(option, " "), ch, int(ducation.Seconds()), file)
}

