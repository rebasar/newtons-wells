package main

import (
	"bytes"
	"fmt"
)

const (
	epsilon float64 = 0.0000001
)

type Polynomial []complex64

type RootResult struct {
	Poly         Polynomial
	Solved       bool
	Steps        int
	InitialGuess complex64
	Root         complex64
}

func (p Polynomial) Derive() Polynomial {
	poly := []complex64(p)
	result := make([]complex64, 0, len(poly)-1)
	for idx, val := range poly {
		if idx > 0 {
			result = append(result, complex(float32(idx), 0)*val)
		}
	}
	return result
}

func (p Polynomial) Evaluate(value complex64) complex64 {
	poly := []complex64(p)
	return Horners(value, poly)
}

func (p Polynomial) IsGoodEnough(guess complex64) bool {
	return mag(p.Evaluate(guess)-p.Evaluate(p.Improve(guess))) <= epsilon
}

func (p Polynomial) Improve(guess complex64) complex64 {
	return guess - (p.Evaluate(guess) / p.Derive().Evaluate(guess))
}

func (p Polynomial) FindRoot(initialGuess complex64) RootResult {
	var tryCount int
	var guess complex64
	for guess, tryCount = initialGuess, 0; ; guess, tryCount = p.Improve(guess), tryCount+1 {
		if tryCount >= maxTries {
			return RootResult{p, false, tryCount, initialGuess, guess}
		}
		if p.IsGoodEnough(guess) {
			return RootResult{p, true, tryCount, initialGuess, guess}
		}
	}
}

func (p Polynomial) String() string {
	poly := []complex64(p)
	var buf bytes.Buffer
	for i := len(poly) - 1; i >= 0; i-- {
		buf.WriteString(fmt.Sprintf("%v", poly[i]))
		buf.WriteString(getXRepr(i))
		if i > 0 {
			buf.WriteString("+")
		}
	}
	return buf.String()
}
