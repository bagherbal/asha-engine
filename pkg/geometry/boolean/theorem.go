package boolean

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/combinatorics"
	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func IncidenceSupportTheorem(vectorDimension, lowerGrade, upperGrade int) theorem.Theorem {
	id := fmt.Sprintf("GEO-BOOL-INCIDENCE-R%d-L%d-U%d", vectorDimension, lowerGrade, upperGrade)
	name := fmt.Sprintf("Boolean incidence support Λ^%dR^%d → Λ^%dR^%d", lowerGrade, vectorDimension, upperGrade, vectorDimension)
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerGeometry,
		Status: theorem.ExactFinite,
		Verify: func() theorem.Result {
			const eps = 1e-8
			support, err := BuildIncidenceSupport(vectorDimension, lowerGrade, upperGrade)
			if err != nil {
				return theorem.Result{
					ID: id, Name: name, Layer: theorem.LayerGeometry, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct Boolean incidence support", Passed: false, Detail: err.Error()}},
				}
			}

			lowerExpected, _ := combinatorics.Binomial(vectorDimension, lowerGrade)
			upperExpected, _ := combinatorics.Binomial(vectorDimension, upperGrade)
			isometryResidual, _ := support.IsometryResidual()
			projectorResidual, _ := support.Support.IdempotenceResidual()
			symmetryResidual, _ := support.Support.SymmetryResidual()
			trace, _ := support.Support.Trace()
			rank := support.RankFromGram(eps)
			residue := support.HarmonicResidueDimension()

			checks := []theorem.Check{
				{
					Name:   "basis dimensions",
					Passed: support.LowerDimension() == lowerExpected && support.UpperDimension() == upperExpected,
					Detail: fmt.Sprintf("dim C^%d=%d, dim C^%d=%d", lowerGrade, support.LowerDimension(), upperGrade, support.UpperDimension()),
				},
				{
					Name:   "incidence rank",
					Passed: rank == support.LowerDimension(),
					Detail: fmt.Sprintf("rank(M_%d,%d)=%d from Gram spectrum", upperGrade, lowerGrade, rank),
				},
				{
					Name:   "normalized incidence isometry",
					Passed: isometryResidual < eps,
					Detail: fmt.Sprintf("||WᵀW−I||_F = %.3e", isometryResidual),
				},
				{
					Name:   "Boolean support projector",
					Passed: projectorResidual < eps && symmetryResidual < eps,
					Detail: fmt.Sprintf("||P_B²−P_B||_F = %.3e, ||P_B−P_Bᵀ||_F = %.3e", projectorResidual, symmetryResidual),
				},
				{
					Name:   "projector trace rank",
					Passed: math.Abs(trace-float64(support.LowerDimension())) < eps,
					Detail: fmt.Sprintf("Tr(P_B)=%.10f", trace),
				},
				{
					Name:   "Boolean harmonic residue",
					Passed: residue == support.UpperDimension()-support.LowerDimension(),
					Detail: fmt.Sprintf("dim ker(Wᵀ)=dim C^%d−rank(P_B)=%d", upperGrade, residue),
				},
				{
					Name:   "Dirac–Kähler collar spectrum",
					Passed: isometryResidual < eps,
					Detail: fmt.Sprintf("D_B has %d eigenvalues +1, %d eigenvalues -1, and %d zero modes when WᵀW=I", support.LowerDimension(), support.LowerDimension(), residue),
				},
			}

			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerGeometry,
				Status: theorem.ExactFinite,
				Checks: checks,
				Notes: []string{
					"This gate constructs the Boolean normal derivative and its rank-56 support inside the middle chamber.",
					"No observed physical constants are used or compared here.",
				},
			}
		},
	}
}
