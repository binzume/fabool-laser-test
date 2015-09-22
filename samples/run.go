package main

import (
	"github.com/binzume/gorecpt1wrap"
	"fmt"
	"time"
)

func main() {
	chs, _ := gorecpt1.GetChannels()
	fmt.Println("Channels:")
	for c := range chs {
		fmt.Println(c)
	}

	cmd := gorecpt1.RecordCmd("1", 1 * time.Minute, "test.ts", []string{"--dev", "/dev/hoge", "--b25", "--strip"})
	fmt.Println("CMD: " + cmd)
}

