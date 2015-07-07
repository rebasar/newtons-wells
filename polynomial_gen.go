package main

func GeneratePolynomial(roots ...complex64) Polynomial {
	rootPolynomials := make([]Polynomial, len(roots), len(roots))
	for i, v := range roots {
		rootPolynomials[i] = toRootPolynomial(v)
	}
	var result Polynomial = rootPolynomials[0]
	for _, v := range rootPolynomials[1:] {
		result = MultiplyPolynomial(result, v)
	}
	return result
}

func toRootPolynomial(c complex64) Polynomial {
	return []complex64{-c, complex(1, 0)}
}
