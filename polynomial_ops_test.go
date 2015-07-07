package main

import (
	"testing"
)

// Shamelessly stolen from https://stackoverflow.com/questions/15311969/checking-the-equality-of-two-slices
func sliceEq(a, b []complex64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func checkSlice(result, expected []complex64, t *testing.T) {
	if !sliceEq(result, expected) {
		t.Errorf("Result \"%v\" is not equal to expected result \"%v\"", result, expected)
	}
}

func TestAddSameLengthPolynomials(t *testing.T) {
	var poly1 Polynomial = []complex64{complex(2, 0), complex(0, 0), complex(3, 0)}
	var poly2 Polynomial = []complex64{complex(0, 0), complex(1, 0), complex(2, 0)}
	var expected Polynomial = []complex64{complex(2, 0), complex(1, 0), complex(5, 0)}
	result := AddPolynomial(poly1, poly2)
	checkSlice(result, expected, t)
}

func TestAddDifferentLengthPolynomials(t *testing.T) {
	var poly1 Polynomial = []complex64{complex(2, 0), complex(0, 0), complex(3, 0)}
	var poly2 Polynomial = []complex64{complex(1, 0), complex(2, 0)}
	var expected Polynomial = []complex64{complex(3, 0), complex(2, 0), complex(3, 0)}
	result := AddPolynomial(poly1, poly2)
	checkSlice(result, expected, t)
}

func TestRaisingPolynomialToDegree(t *testing.T) {
	poly := []complex64{complex(1, 0), complex(2, 0)}
	expected := []complex64{complex(0, 0), complex(1, 0), complex(2, 0)}
	result := raisePolyBy(poly, 1)
	checkSlice(result, expected, t)
}

func TestMultiplyingPolynomialWithCoefficient(t *testing.T) {
	poly := []complex64{complex(1, 0), complex(2, 0)}
	expected := []complex64{complex(0, 0), complex(2, 0), complex(4, 0)}
	result := multiplyWith(poly, 1, complex(2, 0))
	checkSlice(result, expected, t)
}

func TestMultiplySameLengthPolynomials(t *testing.T) {
	var poly1 Polynomial = []complex64{complex(1, 0), complex(1, 0)}
	var poly2 Polynomial = []complex64{complex(-1, 0), complex(1, 0)}
	var expected Polynomial = []complex64{complex(-1, 0), complex(0, 0), complex(1, 0)}
	result := MultiplyPolynomial(poly1, poly2)
	checkSlice(result, expected, t)
}
