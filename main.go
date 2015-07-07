package main

import (
	"image/png"
	"runtime"
)

func main() {
	cmdLine := parseCommandLine()
	if cmdLine.Help {
		usage()
	} else {
		runtime.GOMAXPROCS(min(cmdLine.Concurrency, runtime.NumCPU()))
		cfg := StartingConfig{cmdLine.Roots, 0, 0, cmdLine.Width, cmdLine.Height, cmdLine.Concurrency}
		result := GenerateMatrix(cfg)
		png.Encode(cmdLine.Output, ToImage(cfg, result))
	}
}

// Cannot believe that the Go standard library does not include this
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
