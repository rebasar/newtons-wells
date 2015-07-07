package main

import (
	hsl "code.google.com/p/sadbox/color"
	"fmt"
	"image"
	"image/color"
	"math"
	"math/cmplx"
)

type StartingConfig struct {
	Roots                    []complex64
	Top, Left, Width, Height int
	Concurrency              int
}

type RunningConfig struct {
	Poly               Polynomial
	Begin, End, Column int
}

type PartResult struct {
	Result []RootResult
	Index  int
}

func (cfg *StartingConfig) GetClosestRoot(val complex64) complex64 {
	m := mag(val)
	currentRoot := cfg.Roots[0]
	minDistance := math.Abs(mag(currentRoot) - m)
	for _, v := range cfg.Roots[1:] {
		newDistance := math.Abs(mag(v) - m)
		if newDistance < minDistance {
			minDistance = newDistance
			currentRoot = v
		}
	}
	return currentRoot
}

func (cfg *RunningConfig) Size() int {
	return cfg.End - cfg.Begin
}

func (cfg *RunningConfig) ToIndex(i int) int {
	if i < cfg.Begin || i > cfg.End {
		panic(fmt.Errorf("%d is not a valid index! A valid index must be between %d and %d", i, cfg.Begin, cfg.End))
	}
	return i - cfg.Begin
}

func createGoroutines(cfg StartingConfig) (chan<- RunningConfig, <-chan PartResult) {
	jobs := make(chan RunningConfig, cfg.Width)
	results := make(chan PartResult)
	for i := 0; i < cfg.Concurrency; i++ {
		go func(input <-chan RunningConfig, output chan<- PartResult) {
			for job := range input {
				output <- PartResult{ComputePart(job), job.Column}
			}
		}(jobs, results)
	}
	return jobs, results
}

func submitJobs(jobs chan<- RunningConfig, cfg StartingConfig, poly Polynomial) {
	for i := 0; i < cfg.Width; i++ {
		jobs <- RunningConfig{poly, 0, cfg.Height, i}
	}
	close(jobs)
}

func GenerateMatrix(cfg StartingConfig) [][]RootResult {
	result := make([][]RootResult, cfg.Width, cfg.Width)
	poly := GeneratePolynomial(cfg.Roots...)
	jobs, results := createGoroutines(cfg)
	submitJobs(jobs, cfg, poly)
	for i := 0; i < cfg.Width; i++ {
		r := <-results
		result[r.Index] = r.Result
	}
	return result
}

func ComputePart(cfg RunningConfig) []RootResult {
	result := make([]RootResult, cfg.Size(), cfg.Size())
	for i := cfg.Begin; i < cfg.End; i++ {
		r := cfg.Poly.FindRoot(complex(float32(cfg.Column), float32(i)))
		result[cfg.ToIndex(i)] = r
	}
	return result
}

func Pale(c hsl.HSL, steps int) color.Color {
	ldif := float64(steps) / (2 * float64(maxTries))
	r, g, b := hsl.HSLToRGB(c.H, c.S, c.L-ldif)
	return color.RGBA{r, g, b, 255}
}

func GetRootColor(root complex64) hsl.HSL {
	phase := cmplx.Phase(complex128(root))
	h := phase + math.Pi/(2*math.Pi)
	s := 1.0
	l := 0.5
	return hsl.HSL{H: h, S: s, L: l}
}

func Colorize(cfg StartingConfig, rr RootResult) color.Color {
	root := cfg.GetClosestRoot(rr.Root)
	c := GetRootColor(root)
	return Pale(c, rr.Steps)
}

func ToImage(cfg StartingConfig, mtx [][]RootResult) image.Image {
	img := image.NewRGBA(image.Rect(cfg.Left, cfg.Top, cfg.Width, cfg.Height))
	for i, ra := range mtx {
		for j, rr := range ra {
			if rr.Solved {
				img.Set(i, j, Colorize(cfg, rr))
			} else {
				img.Set(i, j, color.Black)
			}
		}
	}
	return img
}
