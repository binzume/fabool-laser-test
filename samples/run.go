package main

import (
	"github.com/binzume/gorecpt1wrap"
	"fmt"
	"time"
)

func main() {
	rec := gorecpt1.New("/dev/hoge")
	chs, _ := rec.GetChannels()
	fmt.Println("Channels:")
	for c := range chs {
		fmt.Println(c)
	}

	cmd := rec.RecordCmd("1", 1 * time.Minute, "test.ts")
	fmt.Println("CMD: " + cmd)
}

