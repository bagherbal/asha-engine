package clifford

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func StructureTheorem(signature Signature) theorem.Theorem {
	id := fmt.Sprintf("ALG-CLIFFORD-CL-%d-%d", signature.Positive, signature.Negative)
	return theorem.Theorem{
		ID:     id,
		Name:   fmt.Sprintf("Clifford algebra Cℓ(%d,%d)", signature.Positive, signature.Negative),
		Layer:  theorem.LayerAlgebra,
		Status: theorem.ExactFinite,
		Verify: func() theorem.Result {
			alg, err := New(signature)
			if err != nil {
				return theorem.Result{
					ID: id, Name: "Clifford algebra", Layer: theorem.LayerAlgebra, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct Clifford algebra", Passed: false, Detail: err.Error()}},
				}
			}

			n := alg.VectorDimension()
			expected := 1 << n
			checks := []theorem.Check{
				{
					Name:   "algebra dimension",
					Passed: alg.AlgebraDimension() == expected,
					Detail: fmt.Sprintf("dim Cℓ(%d,%d) = 2^%d = %d", signature.Positive, signature.Negative, n, alg.AlgebraDimension()),
				},
				{
					Name:   "metric signature",
					Passed: signature.Positive == 1 && signature.Negative == 7,
					Detail: fmt.Sprintf("diagonal metric = %v", alg.MetricDiagonal()),
				},
				{
					Name:   "time-like generator",
					Passed: alg.AnticommutatorCoefficient(0, 0) == 2,
					Detail: "e₀e₀ + e₀e₀ = +2",
				},
				{
					Name:   "space-like generator",
					Passed: alg.AnticommutatorCoefficient(1, 1) == -2,
					Detail: "e₁e₁ + e₁e₁ = -2",
				},
				{
					Name:   "orthogonal anticommutation",
					Passed: alg.AnticommutatorCoefficient(0, 1) == 0,
					Detail: "e₀e₁ + e₁e₀ = 0 for diagonal metric",
				},
			}

			return theorem.Result{
				ID:     id,
				Name:   fmt.Sprintf("Clifford algebra Cℓ(%d,%d)", signature.Positive, signature.Negative),
				Layer:  theorem.LayerAlgebra,
				Status: theorem.ExactFinite,
				Checks: checks,
				Notes: []string{
					"This verifies the algebraic signature and dimension only.",
					"No physical constant is inferred at this stage.",
				},
			}
		},
	}
}
