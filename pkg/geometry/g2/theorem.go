package g2

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CalibrationSupportTheorem() theorem.Theorem {
	const id = "GEO-G2-CALIBRATION-L4-R8"
	const name = "Octonionic G₂ calibration support inside Λ⁴R⁸"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerGeometry,
		Status: theorem.ExactFinite,
		Verify: func() theorem.Result {
			const eps = 1e-8
			support, err := BuildCalibrationSupport()
			if err != nil {
				return theorem.Result{
					ID: id, Name: name, Layer: theorem.LayerGeometry, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct G₂ calibration support", Passed: false, Detail: err.Error()}},
				}
			}
			isometryResidual, _ := support.IsometryResidual()
			projectorResidual, _ := support.Support.IdempotenceResidual()
			symmetryResidual, _ := support.Support.SymmetryResidual()
			trace, _ := support.Support.Trace()
			rank := support.RankFromGram(eps)

			checks := []theorem.Check{
				{
					Name:   "middle chamber dimension",
					Passed: support.MiddleDimension() == 70,
					Detail: fmt.Sprintf("dim Λ⁴R⁸=%d", support.MiddleDimension()),
				},
				{
					Name:   "associative G₂ form",
					Passed: support.Associative.NonZeroCanonicalTerms() == 7,
					Detail: fmt.Sprintf("φ has %d canonical Fano terms", support.Associative.NonZeroCanonicalTerms()),
				},
				{
					Name:   "coassociative G₂ form",
					Passed: support.Coassociative.NonZeroCanonicalTerms() == 7,
					Detail: fmt.Sprintf("*φ has %d canonical Hodge-dual terms", support.Coassociative.NonZeroCanonicalTerms()),
				},
				{
					Name:   "calibration sector dimension",
					Passed: support.SectorDimension() == 14 && rank == 14,
					Detail: fmt.Sprintf("M₁₄ᴳ = 7_t ⊕ 7_s, rank=%d", rank),
				},
				{
					Name:   "orthonormal calibration frame",
					Passed: isometryResidual < eps,
					Detail: fmt.Sprintf("||Q_GᵀQ_G−I||_F = %.3e", isometryResidual),
				},
				{
					Name:   "G₂ support projector",
					Passed: projectorResidual < eps && symmetryResidual < eps,
					Detail: fmt.Sprintf("||P_G²−P_G||_F = %.3e, ||P_G−P_Gᵀ||_F = %.3e", projectorResidual, symmetryResidual),
				},
				{
					Name:   "projector trace rank",
					Passed: math.Abs(trace-14) < eps,
					Detail: fmt.Sprintf("Tr(P_G)=%.10f", trace),
				},
			}

			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerGeometry,
				Status: theorem.ExactFinite,
				Checks: checks,
				Notes: []string{
					"This gate constructs the octonionic rank-14 matter-calibration sector without using observed physical constants.",
					"The Boolean rank-56 support and the G₂ rank-14 support are still separate objects; their contact is tested in the next gate.",
				},
			}
		},
	}
}
