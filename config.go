package main

import (
	"fmt"
	"github.com/ogier/pflag"
	"io"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type CommandLine struct {
	Help          bool
	Roots         []complex64
	Width, Height int
	Concurrency   int
	Output        io.Writer
}

var maxTries int = 25

func usage() {
	fmt.Printf("Usage: %s [options]:\n\n", os.Args[0])
	pflag.PrintDefaults()
	fmt.Println()
}

func parseCommandLine() CommandLine {
	var roots, outputfile string
	var width, height int
	var concurrency int
	var help bool = false
	pflag.StringVarP(&roots, "roots", "r", "1+1i", "Complex roots, separated by commas. Ex: 3+2i,1+1i")
	pflag.StringVarP(&outputfile, "output", "o", "-", "Output file name. If omitted, standard output will be used.")
	pflag.IntVarP(&width, "width", "x", 640, "Width of the resulting image")
	pflag.IntVarP(&height, "height", "y", 480, "Height of the resulting image")
	pflag.IntVarP(&concurrency, "concurrency", "c", runtime.NumCPU(), "Number of CPU cores to use. Defaults to all the cores available.")
	pflag.IntVarP(&maxTries, "max-tries", "m", maxTries, "Max tries for finding the root")
	pflag.BoolVarP(&help, "help", "h", false, "Usage information")
	pflag.Parse()
	rootValues := parseRoots(roots)
	writer := getWriter(outputfile)
	return CommandLine{help, rootValues, width, height, concurrency, writer}
}

func parseRoots(s string) []complex64 {
	rawValues := strings.Split(s, ",")
	roots := make([]complex64, len(rawValues), len(rawValues))
	for idx, val := range rawValues {
		r, err := parseComplex(strings.TrimSpace(val))
		if err != nil {
			log.Fatalf("Cannot parse \"%v\" as a complex number! Error: %v", val, err)
		}
		roots[idx] = r
	}
	return roots
}

func parseComplex(s string) (complex64, error) {
	exp := regexp.MustCompile("(\\d+)((\\+|-)\\d+)?i?")
	result := exp.FindStringSubmatch(s)
	if result == nil || len(result) == 1 {
		return complex(0, 0), fmt.Errorf("Cannot parse \"%v\" as a complex number!", s)
	} else if len(result) == 2 {
		real, err := strconv.ParseFloat(result[1], 32)
		if err != nil {
			return complex(0, 0), err
		} else {
			return complex(float32(real), 0), nil
		}
	} else {
		real, err1 := strconv.ParseFloat(result[1], 32)
		imag, err2 := strconv.ParseFloat(result[2], 32)
		if err1 != nil {
			return complex(0, 0), err1
		} else if err2 != nil {
			return complex(0, 0), err2
		} else {
			return complex(float32(real), float32(imag)), nil
		}
	}
}

func getWriter(filename string) io.Writer {
	if filename == "-" {
		return os.Stdout
	}
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
