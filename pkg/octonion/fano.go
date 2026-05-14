package octonion

// FanoTerm stores one oriented associative triple for the imaginary octonions.
// Indices 0..6 represent e₁..e₇. The scalar unit is not included here.
type FanoTerm struct {
	I, J, K int
	Sign    int
}

// StandardFanoTerms returns the convention used by the Asha finite-source
// construction:
//
// φ = 123 + 145 + 167 + 246 − 257 − 347 − 356
//
// with zero-based indices internally.
func StandardFanoTerms() []FanoTerm {
	return []FanoTerm{
		{I: 0, J: 1, K: 2, Sign: +1},
		{I: 0, J: 3, K: 4, Sign: +1},
		{I: 0, J: 5, K: 6, Sign: +1},
		{I: 1, J: 3, K: 5, Sign: +1},
		{I: 1, J: 4, K: 6, Sign: -1},
		{I: 2, J: 3, K: 6, Sign: -1},
		{I: 2, J: 4, K: 5, Sign: -1},
	}
}
