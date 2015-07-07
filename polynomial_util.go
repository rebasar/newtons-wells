package main

import (
	"bytes"
	"fmt"
	"math"
)

var (
	sup_digits []rune = []rune{'⁰', '¹', '²', '³', '⁴', '⁵', '⁶', '⁷', '⁸', '⁹'}
)

func Horners(value complex64, poly []complex64) complex64 {
	if len(poly) == 0 {
		return complex(0, 0)
	} else {
		return poly[0] + (value * Horners(value, poly[1:]))
	}
}

func mag(c complex64) float64 {
	return math.Sqrt(math.Pow(float64(real(c)), 2) + math.Pow(float64(imag(c)), 2))
}

func digits(i int) string {
	var result bytes.Buffer
	for b := i; b > 0; b = b / 10 {
		result.WriteRune(sup_digits[i%10])
	}
	return result.String()
}

func getXRepr(i int) string {
	if i == 0 {
		return ""
	} else if i == 1 {
		return "x"
	} else {
		return fmt.Sprintf("x%s", digits(i))
	}
}
