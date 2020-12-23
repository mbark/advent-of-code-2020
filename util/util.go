package util

import (
	"flag"
	"os"
	"runtime/pprof"
	"strings"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func ReadInput(in, splitBy string) []string {
	trimmed := strings.Trim(in, "\n")
	return strings.Split(trimmed, splitBy)
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
