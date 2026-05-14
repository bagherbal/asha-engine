package clifford

type Algebra struct {
	Signature Signature
}

func New(signature Signature) (Algebra, error) {
	if err := signature.Validate(); err != nil {
		return Algebra{}, err
	}
	return Algebra{Signature: signature}, nil
}

func (a Algebra) VectorDimension() int {
	return a.Signature.Dimension()
}

func (a Algebra) AlgebraDimension() int {
	return 1 << a.VectorDimension()
}

func (a Algebra) MetricDiagonal() []int {
	return a.Signature.DiagonalMetric()
}

// AnticommutatorCoefficient returns the scalar coefficient in
// e_i e_j + e_j e_i = coefficient * 1.
//
// For a diagonal Clifford metric this is 2η_ij.
func (a Algebra) AnticommutatorCoefficient(i, j int) int {
	if i != j {
		return 0
	}
	return 2 * a.MetricDiagonal()[i]
}
