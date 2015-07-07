package main

func MultiplyPolynomial(p1, p2 Polynomial) Polynomial {
	poly := []complex64(p1)
	poly2 := []complex64(p2)
	result := make([]complex64, 0, len(p1)+len(p2))
	for i, v := range poly {
		mul := multiplyWith(poly2, i, v)
		result = AddPolynomial(result, mul)
	}
	return result
}

func AddPolynomial(p1, p2 Polynomial) Polynomial {
	poly1 := []complex64(p1)
	poly2 := []complex64(p2)
	result := make([]complex64, max(len(p1), len(p2)))
	for i, _ := range result {
		result[i] = addCoefficients(poly1, poly2, i)
	}
	return result
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}

func addCoefficients(poly1, poly2 []complex64, i int) complex64 {
	if len(poly1) > i && len(poly2) > i {
		return poly1[i] + poly2[i]
	} else if len(poly1) > i {
		return poly1[i]
	} else {
		return poly2[i]
	}
}

func multiplyWith(p []complex64, i int, c complex64) []complex64 {
	result := raisePolyBy(p, i)
	for idx, v := range p {
		result[idx+i] = v * c
	}
	return result
}

func raisePolyBy(p []complex64, pow int) []complex64 {
	result := make([]complex64, len(p)+pow, cap(p)+pow)
	// There must be a better way of doing this with slices.
	// Probably copy(p, result[pow:]). Should try that one.
	for i, v := range p {
		result[i+pow] = v
	}
	return result
}
