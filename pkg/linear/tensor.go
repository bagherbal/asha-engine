package linear

import "fmt"

// Kronecker returns the tensor/Kronecker product a ⊗ b.
func Kronecker(a, b Matrix) Matrix {
	out := NewMatrix(a.Rows()*b.Rows(), a.Cols()*b.Cols())
	for ar := 0; ar < a.Rows(); ar++ {
		for ac := 0; ac < a.Cols(); ac++ {
			scale := a.At(ar, ac)
			for br := 0; br < b.Rows(); br++ {
				for bc := 0; bc < b.Cols(); bc++ {
					out.Set(ar*b.Rows()+br, ac*b.Cols()+bc, scale*b.At(br, bc))
				}
			}
		}
	}
	return out
}

// Commutator returns [a,b] = ab - ba.
func Commutator(a, b Matrix) (Matrix, error) {
	ab, err := a.Mul(b)
	if err != nil {
		return Matrix{}, fmt.Errorf("commutator ab: %w", err)
	}
	ba, err := b.Mul(a)
	if err != nil {
		return Matrix{}, fmt.Errorf("commutator ba: %w", err)
	}
	return ab.Sub(ba)
}
