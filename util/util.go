package util

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"strings"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func ReadInput(in, splitBy string) []string {
	trimmed := strings.Trim(in, "\n")
	return strings.Split(trimmed, splitBy)
}

func WithTime() func() {
	now := time.Now()

	return func() {
		fmt.Printf("time taken: %v\n", time.Now().Sub(now))
	}
}

func WithProfiling() func() {
	flag.Parse()
	if *cpuprofile == "" {
		return func() {}
	}

	f, err := os.Create(*cpuprofile)
	if err != nil {
		panic(err)
	}

	pprof.StartCPUProfile(f)
	return pprof.StopCPUProfile
}
