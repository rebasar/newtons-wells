package main

import (
	"testing"
)

func TestGeneratingSimplePolynomialFromRoots(t *testing.T) {
	var expected Polynomial = []complex64{complex(-1, 0), complex(0, 0), complex(1, 0)}
	result := GeneratePolynomial(complex(1, 0), complex(-1, 0))
	checkSlice(result, expected, t)
}

func TestGeneratingPolynomialFromOneRoot(t *testing.T) {
	var expected Polynomial = []complex64{complex(-4, 0), complex(1, 0)}
	result := toRootPolynomial(complex(4, 0))
	checkSlice(result, expected, t)
}
